package models

import "github.com/jinzhu/gorm"

// Item 历史热榜信息
type Item struct {
	gorm.Model
	Title       string
	Description string
	Thumbnail   string
	URL         string `gorm:"column:url"`
	MD5         string `gorm:"column:md5"`
	Extra       string
	NodeID      int64 `gorm:"column:node_id"` // 所属 node id
}

func init() {
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Item{})
}
