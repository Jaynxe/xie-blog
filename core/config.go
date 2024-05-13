package core

import (
	"log"
	"os"

	"github.com/Jaynxe/xie-blog/config"
	"github.com/Jaynxe/xie-blog/global"
	"gopkg.in/yaml.v3"
)

// 读取配置文件
func InitConfig() {
	const configFile = "config.yaml"
	c := &config.Config{}
	ConfigData, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatal("配置文件读取失败")
	}
	err = yaml.Unmarshal(ConfigData, c)
	if err != nil {
		log.Fatal("配置初始化失败")
	}
	log.Print("配置初始化成功")
	// 全局变量存储配置信息
	global.GVB_CONFIG = c
}
