package models

// LastItem 最新热榜信息
type LastItem struct {
	NodeID int64  `gorm:"column:node_id;primary_key"` // 所属 node id，主键
	Items  string // 最新 Items 信息，json -> string

	Updated int64 // 更新时间
}

func init() {
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&LastItem{})
}
