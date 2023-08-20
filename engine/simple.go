package engine

import (
	"log"

	"discuz-crawler/fetcher"
	"discuz-crawler/model"
	"discuz-crawler/persist"
)

type Simple struct {
	Saver persist.Storage
}

func (e Simple) Run(seeds ...model.Request) {
	if len(seeds) == 0 {
		log.Printf("没有种子数据\n")
		return
	}

	if err := e.Saver.Init(); err != nil {
		log.Printf("初始化存储器失败: %s\n", err.Error())
		return
	}

	var requests []model.Request
	for _, seed := range seeds {
		requests = append(requests, seed)
	}

	count := 0
	for len(requests) > 0 {
		request := requests[0]
		requests = requests[1:]

		parseResult, err := e.worker(request)
		if err != nil {
			continue
		}
		requests = append(requests, parseResult.Requests...)
		e.SaveItems(parseResult.Items, &count)
	}
	e.Saver.Close()
}

func (e Simple) SaveItems(items []interface{}, count *int) {
	for _, item := range items {
		dataItem, ok := item.(model.Video)
		if ok {
			var err error
			dataItem, err = e.Saver.Save(dataItem)
			if err != nil {
				log.Printf("数据 %v 保存出错: %s", item, err)
			} else {
				log.Printf("#%d-item: %d(%s) %s\n", *count, dataItem.Id, dataItem.OutId, dataItem.Title)
			}
			*count++
		}
	}
}

func (e Simple) worker(request model.Request) (model.ParseResult, error) {
	doc, err := fetcher.Fetch(request.Url)
	if err != nil {
		return model.ParseResult{}, err
	}
	return request.ParseFunc(doc, request.Deliver), nil
}
