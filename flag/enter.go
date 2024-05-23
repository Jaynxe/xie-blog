package flag

import "flag"

type Option struct {
	User string
	DB   bool
}

// 
func Parse() Option {
	db := flag.Bool("db", false, "初始化数据库")
	user := flag.String("u", "", "创建用户")

	//解析命令行参数写入注册的flag里
	flag.Parse()
	return Option{
		DB:   *db,
		User: *user,
	}
}

// IsWebStop 是否停止web项目
func IsWebStop(option Option) bool {
	if option.DB {
		return true
	}
	if option.User != "" {
		return true
	}
	return false

}

// SwitchOption 根据命令执行不同的函数
func SwitchOption(option Option) {
	if option.DB {
		Makemigrations()
		return
	}
	if option.User == "admin" || option.User == "user" {
		CreateUser(option.User)
		return
	}
	flag.Usage()
}
