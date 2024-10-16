package model

import (
	"github.com/golang-module/carbon/v2"
)

// 用户模型
type SystemUser struct {
	BaseModel
	Username               string        `gorm:"column:username;uniqueIndex:uidx_username;comment:用户名（工号）" json:"username"`
	ENName                 string        `gorm:"column:enName;not null;comment:英文名（没有就用拼音）" json:"enName"`
	CNName                 string        `gorm:"column:cnName;not null;comment:中文名" json:"cnName"`
	Phone                  string        `gorm:"column:phone;uniqueIndex:uidx_phone;comment:手机号" json:"phone"`
	HidePhone              *uint         `gorm:"column:hidePhone;default:1;comment:是否隐藏手机号(0=不隐藏,1=隐藏)" json:"hidePhone"`
	Email                  string        `gorm:"column:email;uniqueIndex:uidx_email;comment:邮箱" json:"email"`
	Password               string        `gorm:"column:password;not null;comment:密码" json:"-"`
	Secret                 string        `gorm:"column:secret;comment:双因子认证密钥" json:"-"`
	Avatar                 string        `gorm:"column:avatar;default:https://gw.alipayobjects.com/zos/rmsportal/BiazfanxmamNRoxxVxka.png;comment:头像" json:"avatar"`
	Gender                 *uint         `gorm:"column:gender;type:tinyint(1);default:1;comment:性别(1=男,2=女,3=未知)" json:"gender"`
	Birthday               carbon.Carbon `gorm:"column:birthday;comment:生日" json:"birthday"`
	NativePlace            string        `gorm:"column:nativePlace;comment:籍贯" json:"nativePlace"`
	Department             string        `gorm:"column:department;not null;comment:部门" json:"department"`
	JobPosition            string        `gorm:"column:jobPosition;not null;comment:工作岗位" json:"jobPosition"`
	JobId                  string        `gorm:"column:jobId;uniqueIndex:uidx_jobId;comment:工号" json:"jobId"`
	JoinTime               carbon.Carbon `gorm:"column:joinTime;comment:入职日期" json:"joinTime"`
	OfficeAddress          string        `gorm:"column:officeAddress;comment:办公地点详细地址" json:"officeAddress"`
	OfficeStation          string        `gorm:"column:officeStation;comment:工位" json:"officeStation"`
	LastLoginIP            string        `gorm:"column:lastLoginIP;comment:最后一次登录IP" json:"lastLoginIP"`
	LastLoginTime          carbon.Carbon `gorm:"column:lastLoginTime;comment:最后一次登录时间" json:"lastLoginTime"`
	LastChangePasswordTime carbon.Carbon `gorm:"column:lastChangePasswordTime;comment:最后一次修改密码时间" json:"lastChangePasswordTime"`
	Status                 *uint         `gorm:"column:status;type:tinyint(1);default:1;comment:用户状态(0=禁用,1=正常)" json:"status"`
	CreatorId              uint          `gorm:"column:creatorId;comment:创建人id" json:"creatorId"` // 关联用户自身要用指针类型
	Creator                *SystemUser   `gorm:"foreignKey:CreatorId;" json:"creator,omitempty"`
	SystemRoleId           uint          `gorm:"column:systemRoleId;comment:角色id" json:"systemRoleId"` // 关联角色
	SystemRole             SystemRole    `gorm:"foreignKey:SystemRoleId;" json:"systemRole,omitempty"`
}

// 表名设置
func (SystemUser) TableName() string {
	return "system_user"
}
