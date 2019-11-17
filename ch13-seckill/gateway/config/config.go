package config

import (
	"github.com/go-kit/kit/log"
	conf "github.com/keets2012/Micro-Go-Pracrise/ch13-seckill/common/config"
	"github.com/spf13/viper"
	"os"
)

const (
	kConfigType = "CONFIG_TYPE"
)

var Logger log.Logger

func init() {
	Logger = log.NewLogfmtLogger(os.Stderr)
	Logger = log.With(Logger, "ts", log.DefaultTimestampUTC)
	Logger = log.With(Logger, "caller", log.DefaultCaller)
	viper.AutomaticEnv()
	initDefault()

	if err := conf.LoadRemoteConfig(); err != nil {
		Logger.Log("Fail to load remote config", err)
	}
	if err := conf.Sub("auth", &AuthPermitAll{}); err != nil {
		Logger.Log("Fail to parse config", err)
	}
}

func initDefault() {
	viper.SetDefault(kConfigType, "yaml")
}