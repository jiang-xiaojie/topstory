package zhihu

import "testing"

func Test_Craw(t *testing.T) {
	spider := &RebangSpider{}
	spider.Crawl()
	t.Error("hello")
}
