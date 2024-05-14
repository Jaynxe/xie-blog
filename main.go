package main

import (
	"github.com/Jaynxe/xie-blog/core"
	"github.com/Jaynxe/xie-blog/global"
)

func init() {
	core.InitConfig()
	core.InitLogWithLevel("log")
	
	core.InitGorm()
}
func main() {
	global.GVB_LOGGER.Info("hello world")
}
