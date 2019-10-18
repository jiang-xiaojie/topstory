package zhihu

import (
	"fmt"

	"github.com/gocolly/colly"
)

type RebangSpider struct{}

func (spider *RebangSpider) Crawl() {
	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.90 Safari/537.36"

	c.OnHTML("div[class=Card] > a[class=HotList-item]", func(e *colly.HTMLElement) {
		index := e.ChildText("div[class=HotList-itemPre]")
		title := e.ChildText("div[class=HotList-itemBody] > div[class=HotList-itemTitle]")
		url := e.ChildAttr("div[class=HotItem-content] > a", "href")
		extra := e.ChildText("div[class=HotItem-content] > div[class=HotList-itemMetrics]")

		fmt.Println(index, title, extra, url)
	})
	c.Visit("https://www.zhihu.com/billboard")
}
