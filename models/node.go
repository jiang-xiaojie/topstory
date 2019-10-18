package models

import (
	"fmt"

	"github.com/jianggushi/topstory/pkg/utils"
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
	MD5      string `gorm:"column:md5;unique_index"` // 站点在本站的 URL 标识
}

func NewNode(name, display, homepage, logo, domain string) (*Node, error) {
	node := Node{
		Name:     name,
		Display:  display,
		Homepage: homepage,
		Logo:     logo,
		Domain:   domain,
		MD5:      utils.MD5(homepage),
	}
	err := db.Create(&node).Error
	if err != nil {
		return nil, fmt.Errorf("new node: %w", err)
	}
	return &node, nil
}

func GetNodeByHomepage(homepage string) (*Node, error) {
	md5 := utils.MD5(homepage)
	node := Node{}
	err := db.Where("md5 = ?", md5).First(&node).Error
	if err != nil {
		return nil, fmt.Errorf("get node: %w", err)
	}
	return &node, nil
}

func init() {
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Node{})
}
