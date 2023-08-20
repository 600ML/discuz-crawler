package persist

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"discuz-crawler/model"
)

type FileSaver struct {
	File *os.File
}

func (f *FileSaver) Init() error {
	fileName := fmt.Sprintf("%s.txt", time.Now().Format("2006-01-02 15:04:05"))
	var err error
	f.File, err = os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModeAppend)
	if err != nil {
		return fmt.Errorf("创建(打开)文件%s失败: %w", fileName, err)
	}
	return nil
}

func (f *FileSaver) Save(item model.Video) (model.Video, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false) //禁止转义

	if err := encoder.Encode(item); err != nil {
		return item, fmt.Errorf("序列化数据失败: %w", err)
	}

	if _, err := f.File.Write(buffer.Bytes()); err != nil {
		return item, fmt.Errorf("写入数据失败: %w", err)
	}
	return item, nil
}

func (f *FileSaver) Close() {
	_ = f.File.Close()
}
