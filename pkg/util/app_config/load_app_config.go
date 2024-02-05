package app_config

import (
	"strings"

	"github.com/gofrontier-com/pony-express/pkg/core/adf"
	"github.com/spf13/viper"
)

func LoadAppConfig() (appConfig *adf.PonyConfig, err error) {
	viper.SetEnvPrefix("pony")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	err = viper.Unmarshal(&appConfig)

	return
}
