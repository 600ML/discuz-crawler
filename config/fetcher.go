package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type CrawlerConfig struct {
	Storage string `yaml:"storage"`
	Seed    struct {
		Url    string `yaml:"url"`
		Parser string `yaml:"parser"`
	}
	Selector struct {
		Section    string `yaml:"section"`
		SubSection string `yaml:"subSection"`
		NextPage   string `yaml:"next_page"`
		Title      string `yaml:"title"`
		Article    string `yaml:"article"`
	}
	Header map[string]string `yaml:"header"`
}

var Crawler = CrawlerConfig{}

func init() {
	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("读取yaml配置文件config.yaml失败: %s ", err)
	}
	err = yaml.Unmarshal(yamlFile, &Crawler)
	if err != nil {
		log.Fatalf("yaml配置文件格式有误: %s", err)
	}
}
