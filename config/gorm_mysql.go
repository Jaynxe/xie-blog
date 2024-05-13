package config

type Mysql struct {
	Host     string `json:"host,omitempty" yaml:"host"`
	Port     uint   `json:"port,omitempty" yaml:"port"`
	DB       string `json:"db,omitempty" yaml:"db"`
	Username string `json:"username,omitempty" yaml:"username"`
	Password string `json:"password,omitempty" yaml:"password"`
	LogLevel string `json:"log_level,omitempty" yaml:"log-level"`
}
