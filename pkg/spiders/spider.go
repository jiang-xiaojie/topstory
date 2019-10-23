package spiders

import (
	"time"

	log "github.com/sirupsen/logrus"
)

// Spider define a spider interface
type Spider interface {
	Crawl() error
}

var _spiders []Spider

// Register .
func Register(s Spider) {
	_spiders = append(_spiders, s)
}

// RunSpider 定时执行所有的 spider
func RunSpider() {
	log.Info("start run spider")
	runCrawl(time.Now())
	// 每隔 1 分钟执行
	ticker := time.NewTicker(1 * time.Minute)
	for t := range ticker.C {
		runCrawl(t)
	}
}

func runCrawl(t time.Time) {
	log.Infof("run spider time: %v", t)
	for _, spider := range _spiders {
		spider.Crawl()
	}
}
