package config

type Config struct {
	Mysql  Mysql  `yaml:"mysql"`
	Logger Logger `yaml:"logger"`
	System System `yaml:"system"`
	Redis  Redis  `yaml:"redis"`
	Jwt    Jwt    `yaml:"jwt"`
	Local  Local  `yaml:"local"`
	QiNiu  QiNiu  `yaml:"qiniu"`
	Email  Email  `yaml:"email"`
	QQ     QQ     `yaml:"qq"`
	Site   Site   `yaml:"site"`
}
