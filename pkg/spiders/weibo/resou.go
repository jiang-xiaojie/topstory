package weibo

import (
	"errors"
	"strings"

	"github.com/jianggushi/topstory/pkg/spiders"

	"github.com/gocolly/colly"
	"github.com/jianggushi/topstory/models"
	"github.com/jianggushi/topstory/pkg/utils"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

type ResouSpider struct {
	Name     string
	Display  string
	Homepage string
	Logo     string
	Domain   string

	HtmlURL string
	JsonURL string

	NodeID int
}

func (spider *ResouSpider) Crawl() error {
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

func (spider *ResouSpider) parseHTML() ([]*models.Item, error) {
	c := colly.NewCollector()
	items := make([]*models.Item, 0)
	c.OnRequest(func(r *colly.Request) {
		log.Infof("visiting url: %v", r.URL)
	})
	c.OnHTML("tbody > tr", func(e *colly.HTMLElement) {
		if e.ChildText("td[class='td-01 ranktop']") == "" {
			return
		}
		title := e.ChildText("td[class='td-02'] > a")
		URL := e.ChildAttr("td[class='td-02'] > a", "href")
		if strings.HasPrefix(URL, "/") {
			URL = spider.Domain + strings.TrimLeft(URL, "/")
		}
		extra := e.ChildText("td[class='td-02'] > span")
		item := &models.Item{
			Title:  title,
			URL:    URL,
			MD5:    utils.MD5(URL),
			Extra:  extra,
			NodeID: spider.NodeID,
		}
		items = append(items, item)
	})
	c.Visit(spider.HtmlURL)
	if len(items) == 0 {
		return nil, errors.New("not data")
	}
	return items, nil
}

// func (spider *ResouSpider) parseJSON() ([]*models.Item, error) {
// 	c := colly.NewCollector()
// 	items := make([]*models.Item, 0)
// 	c.OnRequest(func(r *colly.Request) {
// 		log.Println("Visiting", r.URL)
// 	})
// 	c.OnResponse(func(r *colly.Response) {
// 		log.Println("response received", r.StatusCode)
// 		js, err := simplejson.NewJson(r.Body)
// 		if err != nil {
// 			return
// 		}
// 		fmt.Println(js.Get("cards").GetIndex(0).Get("card_group"))
// 	})
// 	c.Visit(spider.JsonURL)
// 	if len(items) == 0 {
// 		return nil, errors.New("not data")
// 	}
// 	return items, nil
// }

func init() {
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
			log.Panicf("failed to create new node(%v) error(%v)", spider.Name, err)
		}
	} else if err != nil {
		log.Panicf("failed to get node(%v) error(%v)", spider.Name, err)
	}
	spider.NodeID = int(node.Model.ID)

	spiders.Register(spider)
}
