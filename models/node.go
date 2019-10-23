package models

import (
	"fmt"
	"time"

	"github.com/jianggushi/topstory/pkg/utils"
)

// Node 站点信息
type Node struct {
	ID        int        `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"-"`
	Name      string     `json:"name"`
	Display   string     `json:"display"`
	Homepage  string     `json:"homepage"`
	Logo      string     `json:"logo"`
	Domain    string     `json:"domain"`
	MD5       string     `gorm:"column:md5;unique_index" json:"md5"` // 站点在本站的 URL 标识
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

// GetNodes get nodes
func GetNodes() ([]*Node, error) {
	var nodes []*Node
	err := db.Find(&nodes).Error
	if err != nil {
		return nil, fmt.Errorf("get nodes: %w", err)
	}
	return nodes, nil
}

// GetNodeByID get node by id
func GetNodeByID(id int) (*Node, error) {
	var node Node
	err := db.First(&node, "id = ?", id).Error
	if err != nil {
		return nil, fmt.Errorf("get node by id: %w", err)
	}
	return &node, nil
}

// GetNodeByMD5 get node by md5
func GetNodeByMD5(md5 string) (*Node, error) {
	var node Node
	err := db.First(&node, "md5 = ?", md5).Error
	if err != nil {
		return nil, fmt.Errorf("get node by md5: %w", err)
	}
	return &node, nil
}

// GetNodeByHomepage get node by homepage
func GetNodeByHomepage(homepage string) (*Node, error) {
	var node Node
	md5 := utils.MD5(homepage)
	err := db.First(&node, "md5 = ?", md5).Error
	if err != nil {
		return nil, fmt.Errorf("get node by homepage: %w", err)
	}
	return &node, nil
}

func init() {
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Node{})
}
