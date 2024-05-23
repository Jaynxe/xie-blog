package flag

import (
	"github.com/Jaynxe/xie-blog/global"
	"github.com/Jaynxe/xie-blog/model"
)

// 命令行数据库迁移
func Makemigrations() {

	err := global.GVB_DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.Category{}, &model.User{}, &model.Tag{}, &model.Comment{}, &model.Image{}, &model.MenuItem{})
	if err != nil {
		global.GVB_LOGGER.Fatal("数据库迁移失败")
	}
	global.GVB_LOGGER.Info("数据库迁移成功")
}
