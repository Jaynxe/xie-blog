package image

import (
	"errors"
	"fmt"
	"github.com/Jaynxe/xie-blog/utils/errhandle"

	"github.com/Jaynxe/xie-blog/global"
	"github.com/Jaynxe/xie-blog/model"
	"github.com/Jaynxe/xie-blog/service"
	"github.com/gin-gonic/gin"
)

// UploadFile 文件上传 godoc
// @Summary 文件上传
// @Schemes
// @Description 文件上传
// @Tags image
// @Accept json
// @Produce json
// @Param   Authorization    header    string  true   "登录返回的Token"
// @Param   uploadFile  formData   file  true  "要上传的文件"
// @Success 200 {object} model.CommonResponse[[]model.ImageResponse]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /authrequired/uploadImages [post]
func (i *Image) UploadFile(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		global.GVB_LOGGER.Error("获取表单失败 err:", err.Error())
		model.ThrowError(c, err)
		return
	}

	// 获取所有图片
	files := form.File["files"]
	if files == nil {
		model.Throw(c, errhandle.ParamsError)
		return
	}

	//记录成功上传的文件数
	var count int
	//存放图片上传响应
	var resList []model.ImageResponse
	for _, file := range files {
		serviceRes := service.ServiceApp.ImageService.ImageUpLoadService(file, c)
		//上传失败
		if !serviceRes.IsSucceed {
			resList = append(resList, serviceRes)
			continue
		}
		resList = append(resList, serviceRes)
		count++
	}
	if count == 0 {
		model.ThrowError(c, errors.New("所有图片上传失败"))
	} else {
		model.OKWithMsg(c, resList, fmt.Sprintf("%d张图片上传成功", count))
	}

}
