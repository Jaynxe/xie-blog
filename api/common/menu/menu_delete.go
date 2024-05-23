package menu

import (
	"fmt"

	"github.com/Jaynxe/xie-blog/global"
	"github.com/Jaynxe/xie-blog/model"
	"github.com/gin-gonic/gin"
)

// 删除菜单 godoc
// @Summary 删除菜单
// @Schemes
// @Description 删除菜单
// @Tags menu
// @Accept json
// @Produce json
// @Param   Authorization    header    string  true   "登录返回的Token"
// @Param   DelIdListRequest  body   model.DelIdListRequest  true  "要删除的菜单id列表"
// @Success 200 {object} model.CommonResponse[string]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /authrequired/deleteMenu [delete]
func (m *Menu) MenuDelete(c *gin.Context) {
	var md model.DelIdListRequest
	err := c.ShouldBindJSON(&md)
	if err != nil {
		model.ThrowBindError(c, &md, err)
		return
	}
	var MenuList []model.MenuItem
	//根据id列表查找要删除的数据
	affectedNum := global.GVB_DB.Find(&MenuList, md.IdList).RowsAffected
	if affectedNum == 0 {
		model.ThrowWithMsg(c, "菜单不存在")
		return
	}
	err = global.GVB_DB.Delete(&MenuList).Error
	if err != nil {
		model.ThrowError(c, err)
		return
	}
	model.OK(c, fmt.Sprintf("共%d个菜单删除成功", affectedNum))
}
