package global

import (
	"go.uber.org/zap"

	"gin-element-admin/config"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	GeaDb       *gorm.DB
	GeaRedis    *redis.Client
	GeaConfig   config.Server
	GeaViper    *viper.Viper
	GeaLog      *zap.Logger
)
