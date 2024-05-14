package global

import (
	"github.com/Jaynxe/xie-blog/config"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	GVB_CONFIG *config.Config
	GVB_DB     *gorm.DB
	GVB_LOGGER *logrus.Logger
)
