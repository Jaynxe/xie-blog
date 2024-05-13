package main

import (
	"fmt"

	"github.com/Jaynxe/xie-blog/core"
	"github.com/Jaynxe/xie-blog/global"
)

func main() {
	core.InitConfig()
	fmt.Printf("global.GVB_CONFIG: %v\n", global.GVB_CONFIG)
}