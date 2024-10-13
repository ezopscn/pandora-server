package model

import (
	"gorm.io/gorm"
	"time"
)

// 基础模型
type BaseModel struct {
	Id        uint           `gorm:"primaryKey;column:id;comment:自增编号" json:"id"`
	CreatedAt time.Time      `gorm:"column:createdAt;comment:创建时间" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"column:updatedAt;comment:更新时间" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"column:deletedAt;comment:删除时间" json:"deletedAt"`
}
