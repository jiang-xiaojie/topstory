package zhihu

import (
	"errors"

	"github.com/gocolly/colly"
	"github.com/jianggushi/topstory/models"
	"github.com/jianggushi/topstory/pkg/spiders"
	"github.com/jianggushi/topstory/pkg/utils"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

type RebangSpider struct {
	Name     string
	Display  string
	Homepage string
	Logo     string
	Domain   string

	HtmlURL string
	JsonURL string

	NodeID int
}

func (spider *RebangSpider) Crawl() error {
	log.Infof("crawl: %v %v %v", spider.Name, spider.Display, spider.Homepage)
	items, err := spider.parseHTML()
	if err != nil {
		return err
	}
	err = models.SaveItems(spider.NodeID, items)
	if err != nil {
		return err
	}
	return nil
}

func (spider *RebangSpider) parseHTML() ([]*models.Item, error) {
	c := colly.NewCollector()
	items := make([]*models.Item, 0)
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Add("Cookie", `_zap=09ee8132-fd2b-43d3-9562-9d53a41a4ef5; d_c0="AGDv-acVoQ-PTvS01pG8OiR9v_9niR11ukg=|1561288241"; capsion_ticket="2|1:0|10:1561288248|14:capsion_ticket|44:NjE1ZTMxMjcxYjlhNGJkMjk5OGU4NTRlNDdkZTJhNzk=|7aefc35b3dfd27b74a087dd1d15e7a6bb9bf5c6cdbe8471bc20008feb67e7a9f"; z_c0="2|1:0|10:1561288250|4:z_c0|92:Mi4xeGZsekFBQUFBQUFBWU9fNXB4V2hEeVlBQUFCZ0FsVk5PcXo4WFFBNWFFRnhYX2h0ZFZpWTQ5T3dDMGh5ZTV1bjB3|0cee5ae41ff7053a1e39d96df2450077d37cc9924b337584cf006028b0a02f30"; q_c1=ae65e92b2bbf49e58dee5b2b29e1ffb3|1561288383000|1561288383000; tgw_l7_route=f2979fdd289e2265b2f12e4f4a478330; _xsrf=f8139fd6-b026-4f01-b860-fe219aa63543; tst=h; tshl=`)
		log.Infof("visiting url: %v", r.URL)
	})
	c.OnHTML("div[class='HotList-list'] > section[class='HotItem']", func(e *colly.HTMLElement) {
		title := e.ChildText("div[class='HotItem-content'] > a > h2[class='HotItem-title']")
		description := e.ChildText("div[class='HotItem-content'] > a > p[class='HotItem-excerpt']")
		thumbnail := e.ChildAttr("a[class='HotItem-img'] > img", "src")
		URL := e.ChildAttr("div[class='HotItem-content'] > a", "href")
		extra := e.ChildText("div[class='HotItem-content'] > div")
		log.Info(extra)
		item := &models.Item{
			Title:       title,
			Description: description,
			Thumbnail:   thumbnail,
			URL:         URL,
			MD5:         utils.MD5(URL),
			Extra:       extra,
			NodeID:      spider.NodeID,
		}
		items = append(items, item)
	})
	c.Visit(spider.HtmlURL)
	if len(items) == 0 {
		return nil, errors.New("not data")
	}
	return items, nil
}

func init() {
	spider := &RebangSpider{
		Name:     "知乎",
		Display:  "热榜",
		Homepage: "https://www.zhihu.com/hot",
		Logo:     "",
		Domain:   "https://www.zhihu.com/",
		HtmlURL:  "https://www.zhihu.com/hot",
		JsonURL:  "https://www.zhihu.com/api/v3/feed/topstory/hot-list-wx?limit=50",
	}
	node, err := models.GetNodeByHomepage(spider.Homepage)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		node, err = models.NewNode(spider.Name, spider.Display, spider.Homepage, spider.Logo, spider.Domain)
		if err != nil {
			log.Panicf("failed to create new node(%v) error(%v)", spider.Name, err)
		}
	} else if err != nil {
		log.Panicf("failed to get node(%v) error(%v)", spider.Name, err)
	}
	spider.NodeID = node.ID

	spiders.Register(spider)
}
