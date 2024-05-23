package tag

import (
	"fmt"

	"github.com/Jaynxe/xie-blog/global"
	"github.com/Jaynxe/xie-blog/model"
	"github.com/gin-gonic/gin"
)

// 删除标签 godoc
// @Summary 删除标签
// @Schemes
// @Description 删除标签
// @Tags Tag
// @Accept json
// @Produce json
// @Param   Authorization    header    string  true   "登录返回的Token"
// @Param   DelIdListRequest  body   model.DelIdListRequest  true  "要删除的标签id列表"
// @Success 200 {object} model.CommonResponse[string]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /authrequired/deleteTag [delete]
func (m *Tag) TagDelete(c *gin.Context) {
	var md model.DelIdListRequest
	err := c.ShouldBindJSON(&md)
	if err != nil {
		model.ThrowBindError(c, &md, err)
		return
	}
	var TagList []model.Tag
	//根据id列表查找要删除的数据
	affectedNum := global.GVB_DB.Find(&TagList, md.IdList).RowsAffected
	if affectedNum == 0 {
		model.ThrowWithMsg(c, "标签不存在")
		return
	}
	// TODO:如果有该标签的文章就不给删除
	err = global.GVB_DB.Delete(&TagList).Error
	if err != nil {
		model.ThrowError(c, err)
		return
	}
	model.OK(c, fmt.Sprintf("共%d个标签删除成功", affectedNum))
}
