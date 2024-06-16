package tag

import (
	"fmt"
	"github.com/Jaynxe/xie-blog/utils/errhandle"

	"github.com/Jaynxe/xie-blog/global"
	"github.com/Jaynxe/xie-blog/model"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

// TagUpdate 更新指定标签 godoc
// @Summary 更新指定标签
// @Schemes
// @Description 更新指定标签
// @Tags Tag
// @Accept json
// @Produce json
// @Param   Authorization    header    string  true   "登录返回的Token"
// @Param id path int true "Tag id"
// @Param   TagRequest  body   model.TagRequest  true  "要更新的标签内容"
// @Success 200 {object} model.CommonResponse[string]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /authrequired/updateTag/{id} [patch]
func (t *Tag) TagUpdate(c *gin.Context) {
	id := c.Param("id")
	var mr model.TagRequest
	err := c.ShouldBindJSON(&mr)
	if err != nil {
		model.ThrowBindError(c, &mr, err)
		return
	}
	//查找标签是否存在
	var Tag model.Tag
	count := global.GVB_DB.Limit(1).Find(&Tag, id).RowsAffected
	oldTitle := Tag.Name
	if count == 0 {
		model.Throw(c, errhandle.TagNotExists)
		return
	}

	TagMap := structs.Map(&mr)
	err = global.GVB_DB.Model(&Tag).Updates(TagMap).Error
	if err != nil {
		global.GVB_LOGGER.Error(err.Error())
		model.ThrowError(c, err)
		return
	}
	model.OK(c, fmt.Sprintf("标签[%s]修改为[%s]成功", oldTitle, mr.Name))

}
