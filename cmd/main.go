package main

import (
	"log"

	authsvc "github.com/anazibinurasheed/d-api-gateway/internal/auth-svc"
	configs "github.com/anazibinurasheed/d-api-gateway/internal/config"
	inventorysvc "github.com/anazibinurasheed/d-api-gateway/internal/inventorysvc"
	util "github.com/anazibinurasheed/d-api-gateway/internal/util"
	"github.com/gin-gonic/gin"
)

func main() {
	config, err := configs.LoadConfig()
	if util.HasError(err) {
		log.Fatalln("failed to load config:", err)
	}

	r := gin.New()

	_ = authsvc.RegisterRoutes(r, config.AuthSvcPort)
	inventorysvc.RegisterRoutes(r, config.InventorySvcPort)

	log.Fatalln(r.Run(config.ApiGatewayPort))
}
