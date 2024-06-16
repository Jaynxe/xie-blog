package tag

import (
	"fmt"
	"github.com/Jaynxe/xie-blog/utils/errhandle"

	"github.com/Jaynxe/xie-blog/global"
	"github.com/Jaynxe/xie-blog/model"
	"github.com/gin-gonic/gin"
)

// TagCreate 创建标签 godoc
// @Summary 创建标签
// @Schemes
// @Description 创建标签
// @Tags Tag
// @Accept json
// @Produce json
// @Param   Authorization    header    string  true   "登录返回的Token"
// @Param   TagRequest  body   model.TagRequest  true  "标签的内容"
// @Success 200 {object} model.CommonResponse[[]model.Tag]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /authrequired/addTag [post]
func (t *Tag) TagCreate(c *gin.Context) {
	var mr model.TagRequest
	err := c.ShouldBindJSON(&mr)
	if err != nil {
		model.ThrowBindError(c, &mr, err)
		return
	}
	var TagList []model.Tag
	count := global.GVB_DB.Find(&TagList, "title", mr.Name).RowsAffected
	if count > 0 {
		model.Throw(c, errhandle.TagExists)
		return
	}
	err = global.GVB_DB.Table("tags").Create(&mr).Error
	if err != nil {
		global.GVB_LOGGER.Error("标签添加失败")
		model.ThrowError(c, err)
		return
	}
	model.OKWithMsg(c, "", fmt.Sprintf("添加标签[%s]成功", mr.Name))
}
