package global

import (
	"github.com/cancan927/common-gin/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Logger      *zap.SugaredLogger
	DB          *gorm.DB
	RedisClient *config.RedisClient
)
