package models

import (
	"github.com/jinzhu/gorm"
)

// Node 站点信息
type Node struct {
	gorm.Model
	Name     string
	Display  string
	Homepage string
	Logo     string
	Domain   string
	HashID   string `gorm:"column:hash_id"` // 站点在本站的 URL 标识
}

func init() {
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Node{})
}
