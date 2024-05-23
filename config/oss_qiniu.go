package config

type QiNiu struct {
	Enable        bool    `json:"enable,omitempty" yaml:"enable"` //是否启用七牛云存储
	AccessKey     string  `json:"access_key,omitempty" yaml:"access_key"`
	SecretKey     string  `json:"secret_key,omitempty" yaml:"secret_key"`
	Bucket        string  `json:"bucket,omitempty" yaml:"bucket"` //存储桶名字
	CDN           string  `json:"cdn,omitempty" yaml:"cdn"`       //访问图片的地址前缀
	Zone          string  `json:"zone,omitempty" yaml:"zone"`     //存储的地区
	Size          float64 `json:"size,omitempty" yaml:"size"`     //存储的大小限制
	Prefix        string  `json:"prefix,omitempty" yaml:"prefix"` //上传七牛云的文件夹前缀
	UseHTTPS      bool    `json:"use_https,omitempty" yaml:"use-https"`
	UseCdnDomains bool    `json:"use_cdn_domains,omitempty" yaml:"use-cdn-domains"`
}
