package model

import (
	"os"

	"github.com/Jaynxe/xie-blog/global"
	"github.com/Jaynxe/xie-blog/model/ctype"
	"gorm.io/gorm"
)

// BeforeDelete 删除操作的钩子函数
func (b *Image) BeforeDelete(tx *gorm.DB) (err error) {
	if b.ImageStoreType == ctype.Local {
		//本地图片的删除除了删除数据库中的数据还要删除本地存储
		err := os.Remove(b.URL)
		if err != nil {
			global.GVB_LOGGER.Error(err)
			return err
		}
	}
	return
}
