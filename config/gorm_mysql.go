package config

import "strconv"

type Mysql struct {
	Host     string `json:"host,omitempty" yaml:"host"`
	Port     int   `json:"port,omitempty" yaml:"port"`
	DB       string `json:"db,omitempty" yaml:"db"`
	Config   string `json:"config,omitempty" yaml:"config"`
	Username string `json:"username,omitempty" yaml:"username"`
	Password string `json:"password,omitempty" yaml:"password"`
	LogMode  string `json:"log_level,omitempty" yaml:"log-mode"` //日志模式，可能是日志的级别（如 debug、info、warn、error 等）。
}

func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Host + ":" + strconv.Itoa(m.Port) + ")/" + m.DB + "?" + m.Config
}