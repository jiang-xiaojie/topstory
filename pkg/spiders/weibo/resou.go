package weibo

import (
	"errors"
	"fmt"
	"log"
	"strings"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/gocolly/colly"
	"github.com/jianggushi/topstory/models"
)

type ResouSpider struct {
	nodeID  int64
	htmlURL string
	jsonURL string
	domain  string
}

func (spider *ResouSpider) Crawl() error {
	items, err := spider.parseHTML()
	if err != nil {
		return err
	}
	for _, item := range items {
		log.Println(item)
	}
	return nil
}

func (spider *ResouSpider) saveItems(items []*models.Item) error {
	return nil
}

func (spider *ResouSpider) parseHTML() ([]*models.Item, error) {
	c := colly.NewCollector()
	items := make([]*models.Item, 0)
	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL)
	})
	c.OnHTML("tbody > tr", func(e *colly.HTMLElement) {
		if e.ChildText("td[class='td-01 ranktop']") == "" {
			return
		}
		title := e.ChildText("td[class='td-02'] > a")
		URL := e.ChildAttr("td[class='td-02'] > a", "href")
		if strings.HasPrefix(URL, "/") {
			URL = spider.domain + URL
		}
		extra := e.ChildText("td[class='td-02'] > span")
		item := &models.Item{
			Title: title,
			URL:   URL,
			// MD5:    utils.MD5(URL),
			Extra:  extra,
			NodeID: spider.nodeID,
		}
		items = append(items, item)
	})
	c.Visit(spider.htmlURL)
	if len(items) == 0 {
		return nil, errors.New("not data")
	}
	return items, nil
}

func (spider *ResouSpider) parseJSON() ([]*models.Item, error) {
	c := colly.NewCollector()
	items := make([]*models.Item, 0)
	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL)
	})
	c.OnResponse(func(r *colly.Response) {
		log.Println("response received", r.StatusCode)
		js, err := simplejson.NewJson(r.Body)
		if err != nil {
			return
		}
		fmt.Println(js.Get("cards").GetIndex(0).Get("card_group"))
	})
	c.Visit(spider.jsonURL)
	if len(items) == 0 {
		return nil, errors.New("not data")
	}
	return items, nil
}
