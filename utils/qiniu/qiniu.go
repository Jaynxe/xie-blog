package qiniu

import (
	"bytes"
	"context"
	"errors"
	"fmt"

	"github.com/Jaynxe/xie-blog/config"
	"github.com/Jaynxe/xie-blog/global"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"

	"time"
)

// 获取上传的凭证
func getToken(q config.QiNiu) string {
	assessKey := q.AccessKey
	secretKey := q.SecretKey
	bucket := q.Bucket
	//上传策略
	policy := storage.PutPolicy{
		Scope: bucket,
	}
	//身份验证对象
	mac := qbox.NewMac(assessKey, secretKey)
	//生成上传凭证
	upToken := policy.UploadToken(mac)
	return upToken
}

// 获取上传的配置
func getConf(q config.QiNiu) storage.Config {
	//构建七牛云存储的配置对象
	cfg := storage.Config{}
	//空间对应的机房信息
	zone, _ := storage.GetRegionByID(storage.RegionID(q.Zone))
	cfg.Zone = &zone
	//是否使用https域名
	cfg.UseHTTPS = q.UseHTTPS
	//上传是否使用CDN上传加速
	cfg.UseCdnDomains = q.UseCdnDomains
	return cfg
}

// ImageUpload 上传图片
func ImageUpload(data []byte, imageName string, prefix string) (filePath string, err error) {
	if !global.GVB_CONFIG.QiNiu.Enable {
		return "", errors.New("请启用七牛上传")
	}
	q := global.GVB_CONFIG.QiNiu
	if q.AccessKey == "" || q.SecretKey == "" {
		return "", errors.New("请配置accessKey和secretKey")
	}
	if float64(len(data)/1024/1024) > q.Size {
		return "", errors.New("文件大小超出设定")
	}
	//上传凭证
	token := getToken(q)
	//配置对象
	conf := getConf(q)
	//表单上传器
	uploader := storage.NewFormUploader(&conf)
	//存储七牛云图片上传成功后的返回信息。
	ret := storage.PutRet{}
	//存储额外的上传参数。
	extra := storage.PutExtra{
		Params: map[string]string{},
	}
	dataLen := int64(len(data))
	//获取当前时间
	now := time.Now().Format("20060102150405")
	key := fmt.Sprintf("%s/%s_%s", prefix, now, imageName)
	err = uploader.Put(context.Background(), &ret, token, key, bytes.NewReader(data), dataLen, &extra)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s%s", q.CDN, ret.Key), nil
}
