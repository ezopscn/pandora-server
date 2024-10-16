package data

import (
	"errors"
	"gorm.io/gorm"
	"pandora-server/global"
	"pandora-server/model"
)

// 角色数据
var systemRoles = []model.SystemRole{
	{
		Id:          1,
		Name:        "超级管理员",
		Keyword:     "SuperAdministrator",
		Description: "系统最高权限管理角色，不推荐用户直接加入该角色",
	},
	{
		Id:          2,
		Name:        "集群管理员",
		Keyword:     "ClusterAdministrator",
		Description: "集群管理角色，拥有所有集群的管理权限，但是不具备系统的管理权限",
	},
	{
		Id:          3,
		Name:        "普通用户",
		Keyword:     "Normal",
		Description: "系统普通角色，默认只具备基础的系统权限读写权限",
	},
	{
		Id:          4,
		Name:        "访客",
		Keyword:     "Guest",
		Description: "系统访客角色，默认只具备基础的系统权限的只读权限",
	},
}

// 角色数据初始化
func InitializeSystemRole() {
	for _, item := range systemRoles {
		// 查看数据是否存在，如果不存在才执行创建
		var m model.SystemRole
		err := global.MySQLDB.Where("id = ? OR name = ? OR keyword = ?",
			item.Id,
			item.Name,
			item.Keyword).First(&m).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			global.MySQLDB.Create(&item)
		}
	}
}
