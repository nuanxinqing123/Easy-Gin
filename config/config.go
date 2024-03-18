package config

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"Easy-Gin/config/autoload"
)

type Configuration struct {
	App   autoload.App   `mapstructure:"app" json:"app" yaml:"app"`
	Mysql autoload.Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Redis autoload.Redis `mapstructure:"redis" json:"redis" yaml:"redis"`
	Zap   autoload.Zap   `mapstructure:"zap" json:"zap" yaml:"zap"`
	JWT   autoload.JWT   `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
}

var (
	GinConfig Configuration
	GinDB     *gorm.DB
	GinRedis  *redis.Client
	GinLOG    *zap.Logger
	GinVP     *viper.Viper
)
