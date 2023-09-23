package main

import (
	"log"

	authsvc "github.com/anazibinurasheed/dmart-api-gateway/pkg/auth-svc"
	configs "github.com/anazibinurasheed/dmart-api-gateway/pkg/config"
	inventorysvc "github.com/anazibinurasheed/dmart-api-gateway/pkg/inventorysvc"
	util "github.com/anazibinurasheed/dmart-api-gateway/pkg/util"
	"github.com/gin-gonic/gin"
)

func main() {
	config, err := configs.LoadConfig()
	if util.HasError(err) {
		log.Fatalln("failed to load config:", err)
	}

	r := gin.New()

	authsvc := authsvc.RegisterRoutes(r, config.AuthSvcPort)
	inventorysvc.RegisterRoutes(r, config.InventorySvcPort, authsvc)

	log.Fatalln(r.Run(config.ApiGatewayPort))
}
