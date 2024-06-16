package image

import (
	"fmt"
	"github.com/Jaynxe/xie-blog/utils/errhandle"

	"github.com/Jaynxe/xie-blog/global"
	"github.com/Jaynxe/xie-blog/model"

	"github.com/gin-gonic/gin"
)

// ImageDelete 删除图片 godoc
// @Summary 删除图片
// @Schemes
// @Description 删除图片
// @Tags image
// @Accept json
// @Produce json
// @Param   Authorization    header    string  true   "登录返回的Token"
// @Param   DelIdListRequest  body   model.DelIdListRequest  true  "删除图片id列表"
// @Success 200 {object} model.CommonResponse[string]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /authrequired/deleteImages [delete]
func (i *Image) ImageDelete(c *gin.Context) {
	var mi model.DelIdListRequest
	err := c.ShouldBindJSON(&mi)
	if err != nil {
		model.ThrowBindError(c, &mi, err)
		return
	}
	var imageList []model.Image
	//根据id列表查找要删除的数据
	affectedNum := global.GVB_DB.Find(&imageList, mi.IdList).RowsAffected
	if affectedNum == 0 {
		model.Throw(c, errhandle.ImageNotExists)
		return
	}
	err = global.GVB_DB.Delete(&imageList).Error
	if err != nil {
		model.ThrowError(c, err)
		return
	}
	model.OK(c, fmt.Sprintf("共下·%d个文件删除成功", affectedNum))

}
