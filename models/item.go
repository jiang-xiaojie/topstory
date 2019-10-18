package models

import (
	"encoding/json"

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

func SaveItems(items []*Item) error {
	for _, item := range items {
		if db.Where("md5 = ?", item.MD5).First(item).RecordNotFound() {
			// create new record
		} else {
			// update record
		}
	}
	// save items to lastitem
	data, err := json.Marshal(items)
	if err != nil {
		return err
	}
	return nil
}

func init() {
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Item{})
}
