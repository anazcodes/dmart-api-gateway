package config

import (
	util "github.com/anazibinurasheed/dmart-api-gateway/pkg/util"
	"github.com/spf13/viper"
)

type config struct {
	ApiGatewayPort   string `mapstructure:"API_GATEWAY_PORT"`
	AuthSvcPort      string `mapstructure:"AUTH_SVC_PORT"`
	InventorySvcPort string `mapstructure:"INVENTORY_SVC_PORT"`
	CartSvcPort      string `mapstructure:"CART_SVC_PORT"`
}

func LoadConfig() (cfg config, err error) {
	viper.AddConfigPath("./pkg/config/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if util.HasError(err) {
		return
	}

	viper.Unmarshal(&cfg)
	return
}
