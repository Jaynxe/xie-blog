package imgservice

import (
	"fmt"
	"io"
	"mime/multipart"
	"path"
	"path/filepath"

	"github.com/Jaynxe/xie-blog/global"
	"github.com/Jaynxe/xie-blog/model"
	"github.com/Jaynxe/xie-blog/model/ctype"
	"github.com/Jaynxe/xie-blog/utils/qiniu"
	"github.com/Jaynxe/xie-blog/utils/pwd"
	"github.com/Jaynxe/xie-blog/utils/upload"
	"github.com/gin-gonic/gin"
)

// 允许的文件类型
var allowedFileTypes = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
	".pdf":  true,
	".gif":  true,
	".svg":  true,
	".webp": true,
}
// ImageUpLoadService 文件上传的方法
func (i *ImgService) ImageUpLoadService(file *multipart.FileHeader, c *gin.Context) (res model.ImageResponse) {
	fileName := file.Filename
	filePath := path.Join(global.GVB_CONFIG.Local.Path, file.Filename)
	local := &upload.Local{}
	res.FilePath = filePath
	// 判断图片是否符合白名单
	ext := filepath.Ext(fileName)
	isAllow, ok := allowedFileTypes[ext]
	if !ok || !isAllow {
		res.UploadStatus = fmt.Sprintf("图片类型[%s]不符合白名单要求", ext)
		res.FilePath = ""
		return
	}
	//判断图片大小是否符合设定大小
	size := float64(file.Size) / float64(1024*1024)
	if size >= float64(global.GVB_CONFIG.Local.Size) {
		res.UploadStatus = fmt.Sprintf("图片大小为%.2fMB超于设定大小%dMB", size, global.GVB_CONFIG.Local.Size)
		res.FilePath = ""
		return
	}

	if size >= global.GVB_CONFIG.QiNiu.Size && global.GVB_CONFIG.QiNiu.Enable {
		res.UploadStatus = fmt.Sprintf("图片大小为%.2fMB超于设定大小%dMB", size, int(global.GVB_CONFIG.QiNiu.Size))
		res.FilePath = ""
		return
	}

	// 根据文件内容生成hash值
	fileObj, err := file.Open()
	if err != nil {
		global.GVB_LOGGER.Error(err.Error())
	}
	fileData, _ := io.ReadAll(fileObj)
	hash := pwd.MD5V(fileData)
	// 判断数据库中是否存在该图片
	var image model.Image
	count := global.GVB_DB.Limit(1).Find(&image, "hash = ?", hash).RowsAffected
	if count != 0 {
		res.UploadStatus = "图片已存在"
		res.FilePath = image.URL
		return
	}
	//默认上传到本地
	fileStoreType := ctype.Local
	res.UploadStatus = "图片上传本地成功"
	//是否开启上传七牛云
	if global.GVB_CONFIG.QiNiu.Enable {
		filePath, err = qiniu.ImageUpload(fileData, file.Filename, global.GVB_CONFIG.QiNiu.Prefix)
		if err != nil {
			global.GVB_LOGGER.Error(err.Error())
			res.UploadStatus = err.Error()
			res.FilePath = ""
			return
		}
		res.FilePath = filePath
		res.UploadStatus = "上传七牛云成功"
		res.IsSucceed = true
		fileStoreType = ctype.QiNiuYun

	} else {
		if filePath, _, err = local.UploadFile(file); err != nil {
			global.GVB_LOGGER.Errorf("图片[%s]上传失败err:%s", file.Filename, err.Error())
			res.UploadStatus = err.Error()
			res.FilePath = ""
			return
		}
		res.FilePath = filePath
		res.IsSucceed = true
	}

	// 图片存入数据库
	global.GVB_DB.Create(&model.Image{
		URL:            filePath,
		Name:           fileName,
		Hash:           hash,
		ImageStoreType: fileStoreType,
	})
	return
}
