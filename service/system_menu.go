package service

import (
	"pandora-server/global"
	"pandora-server/model"
)

// 递归生成菜单树
func generateSystemMenuTree(menus []model.SystemMenu, parentId uint) []model.SystemMenu {
	var tree []model.SystemMenu
	for _, menu := range menus {
		if menu.ParentId == parentId {
			menu.Children = generateSystemMenuTree(menus, menu.Id)
			tree = append(tree, menu)
		}
	}
	return tree
}

// 通过角色生成菜单树的数据库方法
func GenerateSystemMenuTreeByRoleIdService(roleId uint) (tree []model.SystemMenu, err error) {
	// 获取菜单数据
	var menus []model.SystemMenu
	// 当角色 Id 传递 0，则查询所有菜单，否则根据角色查询对应菜单
	if roleId == 0 {
		err = global.MySQLDB.Order("sort ASC").Find(&menus).Error
		if err != nil {
			global.SystemLog.Error("查询全部菜单列表失败：", err.Error())
			return
		}
	} else {
		var role model.SystemRole
		err = global.MySQLDB.Where("id = ?", roleId).Preload("SystemMenus").First(&role).Error
		if err != nil {
			global.SystemLog.Error("查询指定角色的菜单列表失败：", err.Error())
			return
		}
		menus = role.SystemMenus
	}

	// 生成菜单树
	tree = generateSystemMenuTree(menus, 0)
	return
}
