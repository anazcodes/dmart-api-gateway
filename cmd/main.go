package main

import (
	"log"

	authsvc "github.com/anazibinurasheed/d-api-gateway/internal/auth-svc"
	configs "github.com/anazibinurasheed/d-api-gateway/internal/config"
	util "github.com/anazibinurasheed/d-api-gateway/internal/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	config, err := configs.LoadConfigs()
	if util.HasError(err) {
		log.Fatalln("failed to load config:", err)
	}

	r := gin.New()

	_ = authsvc.RegisterRoutes(r, config.AuthSvcPort)

	log.Fatalln(r.Run(config.ApiGatewayPort))
}
