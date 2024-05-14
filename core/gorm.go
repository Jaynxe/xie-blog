package core

import (
	"fmt"
	"time"

	"github.com/Jaynxe/xie-blog/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// InitGorm 初始化并返回一个 *gorm.DB 类型的数据库连接对象。
func InitGorm() *gorm.DB {
	if global.GVB_CONFIG.Mysql.Host == "" {
		global.GVB_LOGGER.Fatal("未配置mysql,取消gorm连接")
		return nil
	}
	dsn := global.GVB_CONFIG.Mysql.Dsn()
	// MySQL 日志记录器
	var mysqlLogger logger.Interface

	if global.GVB_CONFIG.System.Env == "info" {
		//显示所有sql
		// logger.Default是一个实现了 logger.Interface 接口的日志记录器对象。
		mysqlLogger = logger.Default.LogMode(logger.Info)
	} else {
		mysqlLogger = logger.Default.LogMode(logger.Error) //只打印错误的sql
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqlLogger,
	})

	if err != nil {
		global.GVB_LOGGER.Fatalf(fmt.Sprintf("[%s]mysql连接失败", dsn))
	}

	sqlDB, _ := db.DB()
	sqlDB.SetConnMaxIdleTime(time.Hour * 4) //连接最大复用时间，不能超过mysql的wait_timeout
	sqlDB.SetMaxIdleConns(10)               //最大空闲连接数
	sqlDB.SetMaxOpenConns(100)              //最多可容纳
	global.GVB_LOGGER.Info("mysql数据库连接成功")
	//返回数据库连接
	return db
}

// func myLogger() logger.Interface {
// 	newLogger := logger.New(
// 		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
// 		logger.Config{
// 			SlowThreshold:        time.Second,   // Slow SQL threshold
// 			LogLevel:             logger.Silent, // 什么都不输出
// 			ParameterizedQueries: true,          // Don't include params in the SQL log
// 			Colorful:             true,          // Disable color
// 		},
// 	)
// 	return newLogger
// }
