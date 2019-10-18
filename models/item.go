package models

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

// Item 历史热榜信息
type Item struct {
	gorm.Model
	Title       string
	Description string
	Thumbnail   string
	URL         string `gorm:"column:url"`
	MD5         string `gorm:"column:md5;unique_index"`
	Extra       string
	NodeID      int64 `gorm:"column:node_id"` // 所属 node id
}

func (item *Item) String() string {
	return fmt.Sprintf("%v - %v - %v", item.Title, item.MD5, item.NodeID)
}

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

func SaveItems(nodeID int64, items []*Item) error {
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
	lastItem := LastItem{NodeID: nodeID, Items: string(data)}
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
