package model

import (
	"time"

	"golang.org/x/net/context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Video struct {
	Id       int64  `json:"id"`
	OutId    string `json:"out-id"`
	Url      string `json:"url"`
	Section  string `json:"section"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	CreateAt string `json:"-"`
	UpdateAt string `json:"-"`
}

func (item Video) TableName() string {
	return "gyyg.video"
}

func Create(ctx context.Context, dbo *gorm.DB, item Video) (Video, error) {
	item.CreateAt = time.Now().Format("2006-01-02 15:04:05")
	item.UpdateAt = item.CreateAt
	err := dbo.WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "out_id"}, {Name: "title"}},
		DoUpdates: clause.AssignmentColumns([]string{"url", "section", "title", "content"}),
	}).Create(&item).Error
	return item, err
}
