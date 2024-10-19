package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
	"pandora-server/dto"
	"pandora-server/global"
	"pandora-server/model"
)

// 通过条件筛选用户列表
func SearchSystemUserListService(ctx *gin.Context) (systemUsers []model.SystemUser, err error) {
	// 解析查询参数
	var req dto.SystemUserSearchRequest
	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		global.SystemLog.Error("用户查询数据解析失败")
		return
	}

	// 查询
	dbt := global.MySQLDB

	// 查询用户名
	if req.Username != "" {
		dbt = dbt.Where("username like ?", "%"+req.Username+"%")
	}

	// 查询中文名
	if req.CNName != "" {
		dbt = dbt.Where("cnName like ?", "%"+req.CNName+"%")
	}

	// 查询英文名
	if req.ENName != "" {
		dbt = dbt.Where("enName like ?", "%"+req.ENName+"%")
	}

	// 查询邮箱
	if req.Email != "" {
		dbt = dbt.Where("email like ?", "%"+req.Email+"%")
	}

	// 查询手机号
	if req.Phone != "" {
		dbt = dbt.Where("phone like ?", "%"+req.Phone+"%")
	}

	// 查询性别
	if req.Gender != nil {
		dbt = dbt.Where("gender = ?", *req.Gender)
	}

	// 查询籍贯
	if req.NativePlace != "" {
		dbt = dbt.Where("nativePlace like ?", "%"+req.NativePlace+"%")
	}

	// 查询部门
	if req.Department != "" {
		dbt = dbt.Where("department like ?", "%"+req.Department+"%")
	}

	// 查询岗位
	if req.JobPosition != "" {
		dbt = dbt.Where("jobPosition like ?", "%"+req.JobPosition+"%")
	}

	// 查询工号
	if req.JobId != "" {
		dbt = dbt.Where("jobId like ?", "%"+req.JobId+"%")
	}

	// 查询办公地址
	if req.OfficeAddress != "" {
		dbt = dbt.Where("officeAddress like ?", "%"+req.OfficeAddress+"%")
	}

	// 查询工位
	if req.OfficeStation != "" {
		dbt = dbt.Where("officeStation like ?", "%"+req.OfficeStation+"%")
	}

	// 查询用户状态
	if req.Status != nil {
		dbt = dbt.Where("status = ?", *req.Status)
	}

	// 查询创建者
	if req.CreatorId != nil {
		dbt = dbt.Where("creatorId = ?", *req.CreatorId)
	}

	// 查询角色
	if req.SystemRoleId != nil {
		dbt = dbt.Where("systemRoleId = ?", *req.SystemRoleId)
	}

	// 最终查询
	err = dbt.Preload(clause.Associations).Find(&systemUsers).Error
	return
}

// 通过条件获取指定用户的详细信息
func SearchSystemUserDetailService(fieldName string, value interface{}) (systemUser model.SystemUser, err error) {
	err = global.MySQLDB.Where(fmt.Sprintf("%s = ?", fieldName), value).Preload(clause.Associations).First(&systemUser).Error
	return
}
