package model

type ResourceBlogs struct {
	ID       int    `json:"id" gorm:"column:id;type:int(11);not null;primary_key"`
	Title    string `json:"title" gorm:"column:title;type:varchar(258);"`         //标题 形式为 title_subtitle
	SubTitle string `json:"sub_title" gorm:"column:sub_title;type:varchar(258);"` //标题 形式为 title_subtitle

	Name        string `json:"name" gorm:"column:name;type:varchar(128);"`
	Author      string `json:"author" gorm:"column:author;type:varchar(128);"`
	Category    string `json:"category" gorm:"column:category;type:varchar(128);"`
	Tags        string `json:"tags" gorm:"column:tags;type:varchar(128);"`
	ContentMd   string `json:"content_md" gorm:"column:content_md;type:text;"`
	ContentHtml string `json:"content_html" gorm:"column:content_html;type:text;"`
	Link        string `json:"link" gorm:"column:links;type:varchar(128);"`
	CreatedAt   int64  `json:"created_at" gorm:"column:created_at;type:bigint(20);"`
	UpdatedAt   int64  `json:"updated_at" gorm:"column:updated_at;type:bigint(20);"`
	DeletedAt   int64  `json:"deleted_at" gorm:"column:deleted_at;type:bigint(20);"`
	Score       int64  `json:"score" gorm:"column:score;type:int(11);"`
}

// TableName 数据库表名
func (ResourceBlogs) TableName() string {
	return "resource_blogs"
}
