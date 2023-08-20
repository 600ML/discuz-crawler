package persist

import "discuz-crawler/model"

type Storage interface {
	Init() error
	Save(item model.Video) (model.Video, error)
	Close()
}
