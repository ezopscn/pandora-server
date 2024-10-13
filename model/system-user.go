package model

import "time"

// 用户模型
type SystemUser struct {
	BaseModel
	Username               string    `gorm:"column:username;uniqueIndex:uidx_username;comment:用户名（工号）" json:"username"`
	ENName                 string    `gorm:"column:enName;not null;comment:英文名（没有就用拼音）" json:"enName"`
	CNName                 string    `gorm:"column:cnName;not null;comment:中文名" json:"cnName"`
	Phone                  string    `gorm:"column:phone;uniqueIndex:uidx_phone;comment:手机号" json:"phone"`
	HidePhone              *uint     `gorm:"column:hidePhone;type:tinyint(1);default:0;comment:是否隐藏手机号(0=不隐藏,1=隐藏)" json:"hidePhone"`
	Email                  string    `gorm:"column:email;uniqueIndex:uidx_email;comment:邮箱" json:"email"`
	Password               string    `gorm:"column:password;not null;comment:密码" json:"-"` // json 中不显示该字段
	Avatar                 string    `gorm:"column:avatar;default:https://gw.alipayobjects.com/zos/rmsportal/KDpgvguMpGfqaHPjicRK.svg;comment:头像" json:"avatar"`
	Gender                 *uint     `gorm:"column:gender;type:tinyint(1);default:1;comment:性别(1=男,2=女,3=未知)" json:"gender"`
	LastLoginIP            string    `gorm:"column:lastLoginIP;comment:最后一次登录IP" json:"lastLoginIP"`
	LastLoginTime          time.Time `gorm:"column:lastLoginTime;comment:最后一次登录时间" json:"lastLoginTime"`
	LastChangePasswordTime time.Time `gorm:"column:lastChangePasswordTime;comment:最后一次修改密码时间" json:"lastChangePasswordTime"`
	Status                 *uint     `gorm:"column:status;type:tinyint(1);default:1;comment:用户状态(0=禁用,1=正常)" json:"status"`
}

// 表名设置
func (SystemUser) TableName() string {
	return "system_user"
}
