package service

import "github.com/Jaynxe/xie-blog/service/imgservice"

type ServiceGroup struct {
	ImageService imgservice.ImgService
}

var ServiceApp = new(ServiceGroup)
