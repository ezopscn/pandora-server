package model

// API 模型
type SystemApi struct {
	Id              uint          `gorm:"column:id;primaryKey;comment:自增编号" json:"id"`
	Api             string        `gorm:"column:api;uniqueIndex:uidx_api;comment:接口URI" json:"api"`
	Method          string        `gorm:"column:method;comment:请求方法" json:"method"`
	Name            string        `gorm:"column:name;uniqueIndex:uidx_name;comment:接口名称" json:"name"`
	Description     string        `gorm:"column:description;comment:接口说明" json:"description"`
	SystemApiTypeId uint          `gorm:"column:systemApiTypeId;comment:接口类型id" json:"systemApiTypeId"`
	SystemApiType   SystemApiType `gorm:"foreignKey:SystemApiTypeId;" json:"systemApiType"`
}

// 表名设置
func (SystemApi) TableName() string {
	return "system_api"
}
