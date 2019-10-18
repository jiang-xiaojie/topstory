package weibo

import (
	"errors"
	"testing"

	"github.com/jinzhu/gorm"

	"github.com/jianggushi/topstory/models"
)

func Test_Craw(t *testing.T) {
	spider := &ResouSpider{
		Name:     "微博",
		Display:  "热搜榜",
		Homepage: "https://s.weibo.com/top/summary?cate=realtimehot",
		Logo:     "",
		Domain:   "https://s.weibo.com/",
		HtmlURL:  "https://s.weibo.com/top/summary?cate=realtimehot",
		JsonURL:  "https://api.weibo.cn/2/guest/page?containerid=106003type=25&t=3&disable_hot=1&filter_type=realtimehot",
	}
	node, err := models.GetNodeByHomepage(spider.Homepage)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		node, err = models.NewNode(spider.Name, spider.Display, spider.Homepage, spider.Logo, spider.Domain)
		if err != nil {
			t.Fatal(err)
		}
	} else if err != nil {
		t.Fatal(err)
	}
	spider.NodeID = int64(node.Model.ID)
	spider.Crawl()
	t.Error("hello")
}
