package main

import (
	"log"

	util "github.com/anazibinurasheed/dmart-api-gateway/internal/util"
	authsvc "github.com/anazibinurasheed/dmart-api-gateway/pkg/auth-svc"
	configs "github.com/anazibinurasheed/dmart-api-gateway/pkg/config"
	inventorysvc "github.com/anazibinurasheed/dmart-api-gateway/pkg/inventorysvc"
	"github.com/gin-gonic/gin"
)

func main() {
	config, err := configs.LoadConfig()
	if util.HasError(err) {
		log.Fatalln("failed to load config:", err)
	}

	r := gin.New()

	authsvc := authsvc.RegisterRoutes(r, config.AuthSvcPort)
	inventorysvc.RegisterRoutes(r, config.InventorySvcPort)

	log.Fatalln(r.Run(config.ApiGatewayPort))
}
