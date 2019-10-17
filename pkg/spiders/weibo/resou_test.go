package weibo

import "testing"

func Test_Craw(t *testing.T) {
	spider := &ResouSpider{
		htmlURL: "https://s.weibo.com/top/summary?cate=realtimehot",
		domain:  "https://s.weibo.com",
		jsonURL: "https://api.weibo.cn/2/guest/page?containerid=106003type=25&t=3&disable_hot=1&filter_type=realtimehot",
	}
	spider.Crawl()
	t.Error("hello")
}
