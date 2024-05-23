package config

type Local struct {
	Size      int    `yaml:"size" json:"size,omitempty"`
	Path      string `mapstructure:"path" json:"path,omitempty" yaml:"path"`                   // 本地文件访问路径
	StorePath string `mapstructure:"store-path" json:"store-path,omitempty" yaml:"store-path"` // 本地文件存储路径
}
