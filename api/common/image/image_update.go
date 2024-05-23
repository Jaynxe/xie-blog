package image
import (
	"fmt"

	"github.com/Jaynxe/xie-blog/global"
	"github.com/Jaynxe/xie-blog/model"
	"github.com/gin-gonic/gin"
)

// 更新图片名称 godoc
// @Summary 更新图片名称
// @Schemes
// @Description 更新图片名称
// @Tags image
// @Accept json
// @Produce json
// @Param   Authorization    header    string  true   "登录返回的Token"
// @Param   ImageUpdateInfo  body   model.ImageUpdateRequest  true  "图片ID和新名称"
// @Success 200 {object} model.CommonResponse[string]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /authrequired/updateImage [patch]
func (i *Image) ImageUpdate(c *gin.Context) {
	var iu model.ImageUpdateRequest
	err := c.ShouldBindJSON(&iu)
	// 参数绑定不正确
	if err != nil {
		model.ThrowBindError(c, &iu, err)
		return
	}
	var Image model.Image
	err = global.GVB_DB.First(&Image, iu.ID).Error
	if err != nil {
		model.ThrowWithMsg(c, "该图片不存在")
		return
	}
	name := Image.Name

	err = global.GVB_DB.Model(&Image).Update("name", iu.Name).Error
	if err != nil {
		model.ThrowWithMsg(c, "图片名称修改失败")
		return
	}
	model.OK(c, fmt.Sprintf("图片名称从%s修改为%s", name, iu.Name))
}
