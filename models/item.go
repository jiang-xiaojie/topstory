package models

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
)

// Item 历史热榜信息
type Item struct {
	ID          int        `gorm:"primary_key" json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `gorm:"index" json:"-"`
	Title       string     `json:"title"`
	Description string     `gorm:"type:text" json:"description"`
	Thumbnail   string     `json:"thumbnail"`
	URL         string     `gorm:"column:url" json:"url"`
	MD5         string     `gorm:"column:md5;unique_index" json:"md5"`
	Extra       string     `json:"extra"`
	NodeID      int        `gorm:"column:node_id" json:"node_id"` // 所属 node id
}

func (item *Item) String() string {
	return fmt.Sprintf("%v - %v - %v", item.Title, item.MD5, item.NodeID)
}

// CreateOrUpdate create or update Item
func (item *Item) CreateOrUpdate() error {
	err := db.Select("id").First(item, "md5 = ?", item.MD5).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			err = db.Create(item).Error
			if err != nil {
				log.Printf("create Item %v: %v", item, err)
				return err
			}
			return nil
		}
		log.Printf("read Item %v: %v", item, err)
		return err
	}
	// update item
	err = db.Model(&item).Update("extra", item.Extra).Error
	if err != nil {
		log.Printf("update Item %v: %v", item, err)
		return err
	}
	return nil
}

// GetItemsByNodeID .
func GetItemsByNodeID(nodeID int) ([]*Item, error) {
	var items []*Item
	err := db.Find(&items, "node_id = ?", nodeID).Error
	if err != nil {
		return nil, fmt.Errorf("get items by nodeID: %w", err)
	}
	return items, nil
}

// SaveItems save Items
func SaveItems(nodeID int, items []*Item) error {
	for _, item := range items {
		err := item.CreateOrUpdate()
		if err != nil {
			log.Printf("create or update Item: %v", err)
		}
	}
	// save items to lastitem
	data, err := json.Marshal(items)
	if err != nil {
		log.Printf("json marshal items: %v", err)
		return err
	}
	lastItem := LastItem{NodeID: nodeID, ItemsText: string(data)}
	err = lastItem.CreateOrUpdate()
	if err != nil {
		log.Printf("create or update LastItem: %v", err)
		return err
	}
	return nil
}

func init() {
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Item{})
}
