package models

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

// LastItem 最新热榜信息
type LastItem struct {
	NodeID    int     `gorm:"column:node_id;primary_key" json:"node_id"` // 所属 node id，主键
	ItemsText string  `gorm:"type:text" json:"-"`                        // 最新 Items 信息，json -> string
	Items     []*Item `gorm:"-" json:"items"`
	Updated   int     `json:"updated"`
}

// CreateOrUpdate create or update LastItem
func (item *LastItem) CreateOrUpdate() error {
	err := db.Select("updated").First(&LastItem{}, "node_id = ?", item.NodeID).Error
	if gorm.IsRecordNotFoundError(err) {
		err = db.Create(item).Error
		if err != nil {
			log.Printf("create LastItem: %v", err)
		}
		return err
	} else if err != nil {
		// return other error
		log.Printf("read LastItem: %v", err)
		return err
	}
	// update item
	err = db.Model(&item).Update(item).Error
	if err != nil {
		log.Printf("update LastItem: %v", err)
	}
	return err
}

// GetLastItemByNodeID .
func GetLastItemByNodeID(nodeID int) (*LastItem, error) {
	var lastItem LastItem
	err := db.First(&lastItem, "node_id = ?", nodeID).Error
	if err != nil {
		return nil, fmt.Errorf("get lastItem by nodeID: %w", err)
	}
	err = json.Unmarshal([]byte(lastItem.ItemsText), &lastItem.Items)
	if err != nil {
		return nil, fmt.Errorf("json unmarshal ItemsText: %w", err)
	}
	return &lastItem, nil
}

func init() {
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&LastItem{})
}
