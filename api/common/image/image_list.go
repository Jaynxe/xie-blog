package image

import (
	"github.com/Jaynxe/xie-blog/global"
	"github.com/Jaynxe/xie-blog/model"
	"github.com/Jaynxe/xie-blog/utils"
	"github.com/gin-gonic/gin"
)

// GetAllImages 获取所有图片 godoc
// @Summary 获取所有图片
// @Schemes
// @Description 获取所有图片
// @Tags image
// @Accept json
// @Produce json
// @Param   Authorization    header    string  true   "登录返回的Token"
// @Success 200 {object} model.CommonResponse[[]model.Image]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /authrequired/getAllImages [get]
func (i *Image) GetAllImages(c *gin.Context) {
	var imageList []model.Image
	if err := global.GVB_DB.Select("id", "url", "name").Find(&imageList).Error; err != nil {
		global.GVB_LOGGER.Error("获取图片列表失败")
		model.ThrowError(c, err)
		return
	}
	model.OKWithMsg(c, imageList, "查询成功")
}

// GetPaginatedImages 分页获取图片 godoc
// @Summary 分页获取图片
// @Schemes
// @Description 分页获取图片
// @Tags image
// @Accept json
// @Produce json
// @Param   Authorization    header    string  true   "登录返回的Token"
// @Param   page    query   int  false  "页码"
// @Param   key    query   string  false  "搜索关键字"
// @Param   limit    query   int  false  "每页大小"
// @Param   sort    query   string  false  "排序规则"
// @Success 200 {object} model.CommonResponse[[]model.Image]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /authrequired/paginatedImages [get]
func (i *Image) GetPaginatedImages(c *gin.Context) {
	var page model.PageRequest
	err := c.ShouldBindQuery(&page)
	if err != nil {
		model.ThrowError(c, err)
		return
	}
	imageList, err := utils.ComList([]model.Image{}, utils.Option{PageRequest: page, Debug: true})
	if err != nil {
		model.ThrowError(c, err)
		return
	}
	model.OKWithMsg(c, imageList, "查询成功")
}
