package core

import (
	"io/fs"
	"log"
	"os"

	"github.com/Jaynxe/xie-blog/config"
	"github.com/Jaynxe/xie-blog/global"
	"gopkg.in/yaml.v3"
)

const configFile = "config.yaml"

// 读取配置文件
func InitConfig() {
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

/* 修改配置文件 */
func UpdateYaml() error {
	byteData, err := yaml.Marshal(global.GVB_CONFIG)
	if err != nil {
		return err
	}
	err = os.WriteFile(configFile, byteData, fs.ModePerm)
	if err != nil {
		return err
	}
	global.GVB_LOGGER.Info("配置文件修改成功")
	return nil
}
