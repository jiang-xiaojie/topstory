package zhihu

import (
	"testing"
)

func Test_parseHTML(t *testing.T) {
	spider := &RebangSpider{
		Name:     "知乎",
		Display:  "热榜",
		Homepage: "https://www.zhihu.com/hot",
		Logo:     "",
		Domain:   "https://www.zhihu.com/",
		HtmlURL:  "https://www.zhihu.com/hot",
		JsonURL:  "https://www.zhihu.com/api/v3/feed/topstory/hot-list-wx?limit=50",
	}
	items, err := spider.parseHTML()
	if err != nil {
		t.Error(err)
	}
	t.Error(items)
}
