package model

import "time"

type ResourceCategory struct {
	ID int `json:"id" gorm:"column:id;type:int(11);not null;primary_key"`
	CategorysName      string    `json:"category_name" gorm:"column:category_name;type:varchar(128);not null"`
	CategorysInfo    string    `json:"category_info" gorm:"column:category_info;type:varchar(128);not null"`

	CategorysLink      string    `json:"category_link" gorm:"column:category_link;type:varchar(128);not null"`
	CategorysCreatedAt time.Time `json:"created_at" gorm:"column:category_created_at;type:bigint(20);not null"`
	CategorysUpdatedAt time.Time `json:"updated_at" gorm:"column:category_updated_at;type:bigint(20);not null"`
	CategorysTags  string    `json:"category_tags" gorm:"column:category_tags;type:varchar(128);not null"`
}

// TableName 数据库表名
func (ResourceCategory) TableName() string {
	return "resource_category"
}
