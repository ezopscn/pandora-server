package global

// Redis Key Prefix
type RedisKeyPrefix struct {
	LoginToken      string // 用户登录 Token 前缀
	LoginWrongTimes string // 用户登录错误次数
}

// 配置 Redis Key Prefix
var RKP = RedisKeyPrefix{
	LoginToken:      "login:token",
	LoginWrongTimes: "login:wrong-times",
}
