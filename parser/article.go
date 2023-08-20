package parser

import (
	"strings"

	"discuz-crawler/config"
	"discuz-crawler/model"
	"github.com/PuerkitoBio/goquery"
)

func ParseArticle(doc *goquery.Document, item model.Video) model.ParseResult {
	parseResult := model.ParseResult{}
	doc.Find(config.Crawler.Selector.Article).Each(func(i int, selection *goquery.Selection) {
		content, _ := selection.Html()
		content = strings.Replace(content, "\n", "", -1)
		//log.Printf("content: %s", content)
		item.Content = content
		parseResult.Items = append(parseResult.Items, item)
	})
	return parseResult
}
