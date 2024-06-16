package menu

import (
	"github.com/Jaynxe/xie-blog/global"
	"github.com/Jaynxe/xie-blog/model"
	"github.com/Jaynxe/xie-blog/utils/errhandle"
	"github.com/gin-gonic/gin"
)

// GetMenu 获取指定菜单 godoc
// @Summary 获取指定菜单
// @Schemes
// @Description 获取指定菜单
// @Tags menu
// @Accept json
// @Produce json
// @Param   Authorization    header    string  true   "登录返回的Token"
// @Param id path int true "Menu id"
// @Success 200 {object} model.CommonResponse[model.MenuItem]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /authrequired/getMenu/{id} [get]
func (m *Menu) GetMenu(c *gin.Context) {
	id := c.Param("id")
	var menu model.MenuItem
	err := global.GVB_DB.Limit(1).Find(&menu, id).Error
	if err != nil {
		model.Throw(c, errhandle.MenuNotExists)
		return
	}
	model.OKWithMsg(c, menu, "菜单查询成功")
}
