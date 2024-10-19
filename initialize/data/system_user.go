package data

import (
	"errors"
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
	"pandora-server/global"
	"pandora-server/model"
	"pandora-server/pkg/trans"
	"pandora-server/pkg/utils"
)

// 用户密码
var defaultPassword = "ezops.cn"

// 用户数据
var systemUsers = []model.SystemUser{
	{
		BaseModel: model.BaseModel{
			Id: 1,
		},
		Username:      "super",
		ENName:        "Super Administrator",
		CNName:        "超管",
		Phone:         "18888888888",
		HidePhone:     trans.Uint(1),
		Email:         "super@ezops.cn",
		Password:      utils.CryptoPassword(defaultPassword),
		Gender:        trans.Uint(1),
		Birthday:      carbon.Now(),
		NativePlace:   "广东深圳",
		Department:    "产品研发中心-技术运维部",
		JobPosition:   "运维开发工程师",
		JobId:         "EZOPS000001",
		JoinTime:      carbon.Now(),
		OfficeAddress: "广东省深圳市福田区运维开发大厦1024层724号",
		OfficeStation: "10-A-24",
		Status:        trans.Uint(1),
		CreatorId:     1,
		SystemRoleId:  1,
	},
	{
		BaseModel: model.BaseModel{
			Id: 2,
		},
		Username:      "cluster",
		ENName:        "Cluster Administrator",
		CNName:        "群管",
		Phone:         "17777777777",
		HidePhone:     trans.Uint(1),
		Email:         "cluster@ezops.cn",
		Password:      utils.CryptoPassword(defaultPassword),
		Gender:        trans.Uint(1),
		Birthday:      carbon.Now(),
		NativePlace:   "广东深圳",
		Department:    "产品研发中心-技术运维部",
		JobPosition:   "高级运维工程师",
		JobId:         "EZOPS000002",
		JoinTime:      carbon.Now(),
		OfficeAddress: "广东省深圳市福田区运维开发大厦1024层724号",
		OfficeStation: "10-A-25",
		Status:        trans.Uint(1),
		CreatorId:     1,
		SystemRoleId:  2,
	},
	{
		BaseModel: model.BaseModel{
			Id: 3,
		},
		Username:      "normal",
		ENName:        "Normal",
		CNName:        "普通",
		Phone:         "16666666666",
		HidePhone:     trans.Uint(1),
		Email:         "normal@ezops.cn",
		Password:      utils.CryptoPassword(defaultPassword),
		Gender:        trans.Uint(1),
		Birthday:      carbon.Now(),
		NativePlace:   "广东深圳",
		Department:    "产品研发中心-技术运维部",
		JobPosition:   "运维工程师",
		JobId:         "EZOPS000003",
		JoinTime:      carbon.Now(),
		OfficeAddress: "广东省深圳市福田区运维开发大厦1024层724号",
		OfficeStation: "10-A-26",
		Status:        trans.Uint(1),
		CreatorId:     1,
		SystemRoleId:  3,
	},
	{
		BaseModel: model.BaseModel{
			Id: 4,
		},
		Username:      "guest",
		ENName:        "Guest",
		CNName:        "访客",
		Phone:         "15555555555",
		HidePhone:     trans.Uint(0),
		Email:         "guest@ezops.cn",
		Password:      utils.CryptoPassword(defaultPassword),
		Gender:        trans.Uint(1),
		Birthday:      carbon.Now(),
		NativePlace:   "广东深圳",
		Department:    "产品研发中心-测试部",
		JobPosition:   "测试工程师",
		JobId:         "EZOPS000004",
		JoinTime:      carbon.Now(),
		OfficeAddress: "广东省深圳市福田区运维开发大厦10层24号",
		OfficeStation: "20-A-48",
		Status:        trans.Uint(0),
		CreatorId:     1,
		SystemRoleId:  4,
	},
}

// 用户数据初始化
func InitializeSystemUser() {
	for _, item := range systemUsers {
		// 查看数据是否存在，如果不存在才执行创建
		var m model.SystemUser
		err := global.MySQLDB.Where("id = ? OR username = ? OR phone = ? OR email = ?",
			item.Id,
			item.Username,
			item.Phone,
			item.Email).First(&m).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			global.MySQLDB.Create(&item)
		}
	}
}
