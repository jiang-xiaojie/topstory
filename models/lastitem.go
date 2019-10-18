package models

import (
	"log"

	"github.com/jinzhu/gorm"
)

// LastItem 最新热榜信息
type LastItem struct {
	NodeID int64  `gorm:"column:node_id;primary_key"` // 所属 node id，主键
	Items  string `gorm:"type:text"`// 最新 Items 信息，json -> string

	Updated int64 // 更新时间
}

func (item *LastItem) CreateOrUpdate() error {
	err := db.Select("updated").First(item, "node_id = ?", item.NodeID).Error
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
	err = db.Model(&item).Update("items", item.Items).Error
	if err != nil {
		log.Printf("update LastItem: %v", err)
	}
	return err
}

func init() {
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&LastItem{})
}
