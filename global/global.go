package global

import (
	"github.com/Jaynxe/xie-blog/config"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	GVB_LOGGER *logrus.Logger
	GVB_CONFIG *config.Config
	GVB_DB     *gorm.DB
	GVB_REDIS  *redis.Client
)
