package model

// API 类型模型
type SystemApiType struct {
	Id         uint        `gorm:"column:id;primaryKey;comment:自增编号" json:"id"`
	Name       string      `gorm:"column:name;uniqueIndex:uidx_name;comment:接口类型名称" json:"name"`
	SystemApis []SystemApi `gorm:"-" json:"systemApis,omitempty"`
}

// 表名设置
func (SystemApiType) TableName() string {
	return "system_api_type"
}
