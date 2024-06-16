package site

import (
	"github.com/Jaynxe/xie-blog/global"
	"github.com/Jaynxe/xie-blog/model"
	"github.com/gin-gonic/gin"
)

// GetSiteInfo 获取站点信息 godoc
// @Summary 获取站点信息
// @Schemes
// @Description 获取站点信息
// @Tags site
// @Accept json
// @Produce json
// @Param   Authorization    header    string  true   "登录返回的Token"
// @Success 200 {object} model.CommonResponse[any]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /authrequired/getSiteInfo [get]
func (s *Site) GetSiteInfo(c *gin.Context) {
	model.OK(c, global.GVB_CONFIG.Site)
}
