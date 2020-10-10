package model

import "time"

type ResourceDocs struct {
	ID int `json:"id" gorm:"column:id;type:int(11);not null;primary_key"`
	DocName      string    `json:"doc_name" gorm:"column:doc_name;type:varchar(128);not null"`
	DocAuthor    string    `json:"doc_author" gorm:"column:doc_author;type:varchar(128);not null"`
	DocCategory  string    `json:"doc_category" gorm:"column:doc_category;type:varchar(128);not null"`
	DocTags  string    `json:"doc_tags" gorm:"column:doc_tags;type:varchar(128);not null"`
	DocContent   string    `json:"doc_content" gorm:"column:doc_content;type:text;not null"`
	DocLink      string    `json:"doc_link" gorm:"column:doc_link;type:varchar(128);not null"`
	DocCreatedAt time.Time `json:"created_at" gorm:"column:doc_created_at;type:bigint(20);not null"`
	DocUpdatedAt time.Time `json:"updated_at" gorm:"column:doc_updated_at;type:bigint(20);not null"`
}

// TableName 数据库表名
func (ResourceDocs) TableName() string {
	return "resource_docs"
}
