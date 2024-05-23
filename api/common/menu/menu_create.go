package menu
import (
	"fmt"

	"github.com/Jaynxe/xie-blog/global"
	"github.com/Jaynxe/xie-blog/model"
	"github.com/gin-gonic/gin"
)

// 创建菜单 godoc
// @Summary 创建菜单
// @Schemes
// @Description 创建菜单
// @Tags menu
// @Accept json
// @Produce json
// @Param   Authorization    header    string  true   "登录返回的Token"
// @Param   MenuRequest  body   model.MenuRequest  true  "菜单的内容"
// @Success 200 {object} model.CommonResponse[[]model.MenuItem]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /authrequired/addMenu [post]
func (m *Menu) MenuCreate(c *gin.Context) {
	var mr model.MenuRequest
	err := c.ShouldBindJSON(&mr)
	if err != nil {
		model.ThrowBindError(c, &mr, err)
		return
	}
	var menuList []model.MenuItem
	count := global.GVB_DB.Find(&menuList, "title = ? and url = ?", mr.Title, mr.URL).RowsAffected
	if count > 0 {
		model.ThrowWithMsg(c, fmt.Sprintf("菜单[%s]已存在", mr.Title))
		return
	}
	err = global.GVB_DB.Table("menu_items").Create(&mr).Error
	if err != nil {
		global.GVB_LOGGER.Error("菜单添加失败")
		model.ThrowError(c, err)
		return
	}
	model.OKWithMsg(c, "", fmt.Sprintf("添加菜单[%s]成功", mr.Title))
}
