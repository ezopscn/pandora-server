package middleware

import (
	"errors"
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm/clause"
	"pandora-server/dto"
	"pandora-server/global"
	"pandora-server/model"
	"pandora-server/pkg/gedis"
	"pandora-server/pkg/response"
	"pandora-server/pkg/trans"
	"pandora-server/pkg/utils"
	"time"
)

// JWT 认证中间件
func JWTAuth() (*jwt.GinJWTMiddleware, error) {
	return jwt.New(&jwt.GinJWTMiddleware{
		// 定义一个域，用于展示给用户。一般作为 HTTP 认证时的提示信息
		Realm: global.Config.JWT.Realm,
		// 签名 JWT 的密钥，用于加密和解密令牌的字符串
		Key: []byte(global.Config.JWT.Key),
		// Token 有效期
		Timeout: time.Duration(global.Config.JWT.Timeout) * time.Second,
		// 用户登录校验
		Authenticator: authenticator,
		// Token 封装
		PayloadFunc: payloadFunc,
		// 登录成功响应
		LoginResponse: loginResponse,
		// 登录，认证失败响应
		Unauthorized: unauthorized,
		// 解析 Token
		IdentityHandler: identityHandler,
		// 验证 Token
		Authorizator: authorizator,
		// 注销登录
		LogoutResponse: logoutResponse,
		// Token 查找的字段
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// Token 请求头名称
		TokenHeadName: "Bearer",
	})
}

// 隶属 Login 中间件，当调用 LoginHandler 就会触发
// 通过从 ctx 中检索出数据，进行用户登录认证
// 返回包含用户信息的 Map 或者 Struct
func authenticator(ctx *gin.Context) (interface{}, error) {
	// 1.获取用户登录提交的数据
	var req dto.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return nil, errors.New("获取用户登录信息失败")
	}

	// 2.获取客户端 IP，确保代理透传客户端真实 IP，如果获取 IP 失败则使用 None 做标识
	ip := ctx.ClientIP()
	if ip == "" {
		ip = "None"
	}

	// 3.判断错误登录 Key 的值是否已经到达登录上限
	// 通过 Redis 保存用户登录次数 ，为了避免恶意登录，导致正常用户都无法登录，则 Key 需要包含登录账号和客户端 IP 标识
	key := fmt.Sprintf("%s:%s:%s", global.RKP.LoginWrongTimes, ip, req.Account)
	conn := gedis.NewRedisConnection()
	loginWrongTimes := conn.GetInt(key).UnwrapWithDefaultValue(0)
	if loginWrongTimes >= global.Config.Login.WrongTimes {
		return nil, errors.New("用户登录错误次数已经达到上限")
	}

	// 4.账户密码验证
	var systemUser model.SystemUser
	if utils.IsEmail(req.Account) {
		global.MySQLDB.Where("email = ?", req.Account).Preload(clause.Associations).First(&systemUser)
	} else if utils.IsPhoneNumber(req.Account) {
		global.MySQLDB.Where("phone = ?", req.Account).Preload(clause.Associations).First(&systemUser)
	} else {
		global.MySQLDB.Where("username = ?", req.Account).Preload(clause.Associations).First(&systemUser)
	}
	if !utils.ComparePassword(systemUser.Password, req.Password) {
		conn.Set(key, loginWrongTimes+1, gedis.WithExpire(time.Duration(global.Config.Login.LockTime)*time.Second))
		return nil, errors.New("账户名密码错误")
	}

	// 5.验证用户状态是否正常
	if systemUser.Status == trans.Uint(0) {
		return nil, errors.New("用户已禁用，请联系管理员")
	}

	// 6.登录成功，则删除登录失败的 Redis Key，更新用户登录信息
	_, _ = conn.Del(key)
	global.MySQLDB.Model(&model.SystemUser{}).Where("id = ?", systemUser.Id).Updates(map[string]interface{}{
		"lastLoginIP":   ip,
		"lastLoginTime": carbon.Now(),
	})

	// 7.返回用户登录数据给 PayloadFunc 使用
	ctx.Set("username", systemUser.Username)
	return &systemUser, nil
}

// 隶属 Login 中间件，接收 Authenticator 验证成功后传递过来的数据，进行封装成 Token
// MapClaims 必须包含 IdentityKey
// MapClaims 会被嵌入 Token 中，后续可以通过 ExtractClaims 对 Token 进行解析获取到
func payloadFunc(data interface{}) jwt.MapClaims {
	// 断言判断获取传递过来数据是不是用户数据
	if systemUser, ok := data.(*model.SystemUser); ok {
		// 封装一些常用的字段，方便直接使用前端和后端都能直接使用
		return jwt.MapClaims{
			jwt.IdentityKey:     systemUser.Id,
			"username":          systemUser.Username,
			"cnName":            systemUser.CNName,
			"enNName":           systemUser.ENName,
			"phone":             systemUser.Phone,
			"email":             systemUser.Email,
			"avatar":            systemUser.Avatar,
			"department":        systemUser.Department,
			"jobPosition":       systemUser.JobPosition,
			"systemRoleId":      systemUser.SystemRoleId,
			"systemRoleKeyword": systemUser.SystemRole.Keyword,
			"systemRoleName":    systemUser.SystemRole.Name,
		}
	}
	return jwt.MapClaims{}
}

// 隶属 Login 中间件，响应用户请求
// 接收 PayloadFunc 传递过来的 Token 信息，返回登录成功
func loginResponse(ctx *gin.Context, code int, token string, expire time.Time) {
	// 登录响应数据
	var resp dto.LoginResponse
	resp.Token = token
	resp.Expire = expire.Format(global.SecondTimeFormat)

	// 判断是否允许用户多设备登录
	if !global.Config.Login.MultiDevices {
		v, _ := ctx.Get("username")
		username, ok := v.(string)
		if !ok || username == "" {
			response.FailedWithMessage("用户登录失败")
			return
		}

		// 将 Token 保存到 Redis 中
		key := global.RKP.LoginToken + ":" + username
		conn := gedis.NewRedisConnection()
		conn.Set(key, token, gedis.WithExpire(time.Duration(global.Config.JWT.Timeout)*time.Second))
	}
	response.SuccessWithData(resp)
}

// 登录失败，验证失败的响应
func unauthorized(ctx *gin.Context, code int, message string) {
	response.FailedWithCodeAndMessage(response.RequestUnauthorized, message)
}

// 用户登录后的中间件，用于解析 Token
func identityHandler(ctx *gin.Context) interface{} {
	// 从 Context 中获取用户名
	claims := jwt.ExtractClaims(ctx)
	username, _ := claims["username"].(string)
	return &model.SystemUser{
		Username: username,
	}
}

// 用户登录后的中间件，用于验证 Token
func authorizator(data interface{}, ctx *gin.Context) bool {
	systemUser, ok := data.(*model.SystemUser)
	if ok {
		// 不允许多设备登录配置，验证 Redis 中的数据是否一致
		if !global.Config.Login.MultiDevices {
			token := jwt.GetToken(ctx)
			key := global.RKP.LoginToken + ":" + systemUser.Username
			conn := gedis.NewRedisConnection()
			if conn.GetString(key).Unwrap() != token {
				return false
			}
		}
		return true
	}
	return false
}

// 注销登录
func logoutResponse(ctx *gin.Context, code int) {
	claims := jwt.ExtractClaims(ctx)
	username, _ := claims["username"].(string)
	key := global.RKP.LoginToken + ":" + username
	conn := gedis.NewRedisConnection()
	_, _ = conn.Del(key)
	response.Success()
}
