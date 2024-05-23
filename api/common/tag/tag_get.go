package tag

import (
	"github.com/Jaynxe/xie-blog/global"
	"github.com/Jaynxe/xie-blog/model"
	"github.com/Jaynxe/xie-blog/utils"
	"github.com/gin-gonic/gin"
)

// 获取指定标签 godoc
// @Summary 获取指定标签
// @Schemes
// @Description 获取指定标签
// @Tags Tag
// @Accept json
// @Produce json
// @Param id path int true "Tag id"
// @Param   Authorization    header    string  true   "登录返回的Token"
// @Success 200 {object} model.CommonResponse[model.Tag]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /authrequired/getTag/{id} [get]
func (m *Tag) GetTag(c *gin.Context) {
	id := c.Param("id")
	var Tag model.Tag
	err := global.GVB_DB.Limit(1).Find(&Tag, id).Error
	if err != nil {
		model.ThrowWithMsg(c, "标签不存在")
		return
	}
	model.OKWithMsg(c, Tag, "标签查询成功")
}

// 获取所有标签 godoc
// @Summary 获取所有标签
// @Schemes
// @Description 获取所有标签
// @Tags Tag
// @Accept json
// @Produce json
// @Param   Authorization    header    string  true   "登录返回的Token"
// @Success 200 {object} model.CommonResponse[[]model.Tag]
// @Failure 400  {object} model.CommonResponse[any]
// @Router	/authrequired/getAllTags [get]
func (m *Tag) GetAllTags(c *gin.Context) {
	var ml []model.Tag
	err := global.GVB_DB.Find(&ml).Error
	if err != nil {
		global.GVB_LOGGER.Error("标签查询失败")
		model.ThrowError(c, err)
		return
	}
	model.OKWithMsg(c, ml, "标签查询成功")
}

// 分页获取标签 godoc
// @Summary 分页获取标签
// @Schemes
// @Description 分页获取标签
// @Tags Tag
// @Accept json
// @Produce json
// @Param   Authorization    header    string  true   "登录返回的Token"
// @Param   page    query   int  false  "页码"
// @Param   key    query   string  false  "搜索关键字"
// @Param   limit    query   int  false  "每页大小"
// @Param   sort    query   string  false  "排序规则"
// @Success 200 {object} model.CommonResponse[[]model.Tag]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /authrequired/paginatedTags [get]
func (i *Tag) GetPaginatedTags(c *gin.Context) {
	var page model.PageRequest
	err := c.ShouldBindQuery(&page)
	if err != nil {
		model.ThrowError(c, err)
		return
	}
	TagList, err := utils.ComList([]model.Tag{}, utils.Option{PageRequest: page, Debug: true})
	if err != nil {
		model.ThrowError(c, err)
		return
	}
	model.OKWithMsg(c, TagList, "查询成功")
}
