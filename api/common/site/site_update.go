package site

import (
	"encoding/json"

	"github.com/Jaynxe/xie-blog/config"
	"github.com/Jaynxe/xie-blog/core"
	"github.com/Jaynxe/xie-blog/global"
	"github.com/Jaynxe/xie-blog/model"
	"github.com/Jaynxe/xie-blog/utils"
	"github.com/gin-gonic/gin"
)

// UpdateSiteInfo 修改站点信息 godoc
// @Summary 修改站点信息
// @Schemes
// @Description 修改站点信息
// @Tags site
// @Accept json
// @Produce json
// @Param   Authorization    header    string  true   "登录返回的Token"
// @Param   SiteInfo    body    config.Site  true   "要修改的站点信息"
// @Success 200 {object} model.CommonResponse[string]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /authrequired/updateSiteInfo [patch]
func (s *Site) UpdateSiteInfo(c *gin.Context) {
	b, err := c.GetRawData()
	if err != nil {
		model.ThrowError(c, err)
		return
	}
	var cs config.Site
	err = json.Unmarshal(b, &cs)
	if err != nil {
		model.ThrowError(c, err)
		return
	}
	utils.IgnoreStructCopy(&global.GVB_CONFIG.Site, cs, "")
	err = core.UpdateYaml()
	if err != nil {
		model.ThrowError(c, err)
		return
	}
	model.OK(c, "修改站点信息成功")
}
