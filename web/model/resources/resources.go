package model

import "time"

type ResourceArticle struct {
	ID int `json:"id" gorm:"column:id;type:int(11);not null;primary_key"`

	ArticleName      string    `json:"article_name" gorm:"column:article_name;type:varchar(128);not null"`
	ArticleAuthor    string    `json:"article_author" gorm:"column:article_author;type:varchar(128);not null"`
	ArticleCategory  string    `json:"article_category" gorm:"column:article_category;type:varchar(128);not null"`
	ArticleContent   string    `json:"article_content" gorm:"column:article_content;type:text;not null"`
	ArticleLink      string    `json:"article_link" gorm:"column:article_link;type:varchar(128);not null"`
	ArticleCreatedAt time.Time `json:"created_at" gorm:"column:article_created_at;type:bigint(20);not null"`
	ArticleUpdatedAt time.Time `json:"updated_at" gorm:"column:article_updated_at;type:bigint(20);not null"`
}

// TableName 数据库表名
func (ResourceArticle) TableName() string {
	return "resource_articles"
}
