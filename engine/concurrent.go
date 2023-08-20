package engine

import (
	"log"

	"discuz-crawler/fetcher"
	"discuz-crawler/model"
	"discuz-crawler/persist"
)

type Concurrent struct {
	Saver       persist.Storage
	WorkerCount int
}

func (e *Concurrent) Run(seeds ...model.Request) {
	if len(seeds) == 0 {
		log.Printf("没有种子数据\n")
		return
	}

	in := make(chan model.Request, 10000)
	out := make(chan model.ParseResult, 10000)

	if err := e.Saver.Init(); err != nil {
		log.Printf("初始化存储器失败: %s\n", err.Error())
		return
	}

	workerNum := 0
	for i := 0; i < e.WorkerCount; i++ {
		workerNum++
		createWorker(in, out)
	}

	for _, request := range seeds {
		in <- request
	}

	count := 0
	for {
		result := <-out
		for _, request := range result.Requests {
			in <- request
		}

		e.SaveItems(result.Items, &count)
	}
}

func (e *Concurrent) SaveItems(items []interface{}, count *int) {
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

func Worker(request model.Request) (model.ParseResult, error) {
	doc, err := fetcher.Fetch(request.Url)
	if err != nil {
		return model.ParseResult{}, err
	}
	return request.ParseFunc(doc, request.Deliver), nil
}

func createWorker(in chan model.Request, out chan model.ParseResult) {
	go func() {
		for {
			request := <-in
			result, err := Worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
