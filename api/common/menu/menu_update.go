package menu

import (
	"fmt"
	"github.com/Jaynxe/xie-blog/utils/errhandle"

	"github.com/Jaynxe/xie-blog/global"
	"github.com/Jaynxe/xie-blog/model"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

// MenuUpdate 更新指定菜单 godoc
// @Summary 更新指定菜单
// @Schemes
// @Description 更新指定菜单
// @Tags menu
// @Accept json
// @Produce json
// @Param   Authorization    header    string  true   "登录返回的Token"
// @Param id path int true "Menu id"
// @Param   MenuRequest  body   model.MenuRequest  true  "要更新的菜单内容"
// @Success 200 {object} model.CommonResponse[string]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /authrequired/updateMenu/{id} [patch]
func (m *Menu) MenuUpdate(c *gin.Context) {
	id := c.Param("id")
	var mr model.MenuRequest
	err := c.ShouldBindJSON(&mr)
	if err != nil {
		model.ThrowBindError(c, &mr, err)
		return
	}
	//查找菜单是否存在
	var menu model.MenuItem
	count := global.GVB_DB.Limit(1).Find(&menu, id).RowsAffected
	oldTitle := menu.Title
	if count == 0 {
		model.Throw(c, errhandle.MenuNotExists)
		return
	}

	menuMap := structs.Map(&mr)
	err = global.GVB_DB.Model(&menu).Updates(menuMap).Error
	if err != nil {
		global.GVB_LOGGER.Error(err.Error())
		model.ThrowError(c, err)
		return
	}
	model.OK(c, fmt.Sprintf("菜单[%s]修改为[%s]成功", oldTitle, mr.Title))

}
