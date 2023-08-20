package main

import (
	"discuz-crawler/config"
	"discuz-crawler/engine"
	"discuz-crawler/model"
	"discuz-crawler/parser"
	"discuz-crawler/persist"
)

func main() {
	//e := engine.Simple{
	//	Saver: &persist.MysqlSaver{},
	//}
	e := engine.Concurrent{
		Saver:       &persist.MysqlSaver{},
		WorkerCount: 100,
	}
	e.Run(model.Request{
		Url:       config.Crawler.Seed.Url,
		ParseFunc: parser.StrToFuncOfParser(config.Crawler.Seed.Parser),
	})
}
