package ctype

import "encoding/json"

type ImageStoreType int

const (
	Local    ImageStoreType = 1 //存储在本地图片
	QiNiuYun ImageStoreType = 2 //存储在七牛云的图片

)

func (s ImageStoreType) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}
func (s ImageStoreType) String() string {
	var str string
	switch s {
	case Local:
		str = "本地"
	case QiNiuYun:
		str = "七牛云"
	default:
		str = "其他"
	}
	return str
}