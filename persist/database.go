package persist

import (
	"context"
	"fmt"

	"discuz-crawler/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlSaver struct {
	Dbo *gorm.DB
}

func (m *MysqlSaver) Init() error {
	dbo, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/gyyg?charset=utf8mb4"), &gorm.Config{})
	if err != nil {
		return err
	}
	m.Dbo = dbo
	return nil
}

func (m *MysqlSaver) Save(item model.Video) (model.Video, error) {
	var err error
	item, err = model.Create(context.Background(), m.Dbo, item)
	if err != nil {
		err = fmt.Errorf("写入数据失败: %w", err)
	}
	return item, err
}

func (m *MysqlSaver) Close() {
}
