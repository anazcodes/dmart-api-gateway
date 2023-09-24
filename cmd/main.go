package main

import (
	_ "github.com/anazibinurasheed/dmart-api-gateway/api/docs"
	authsvc "github.com/anazibinurasheed/dmart-api-gateway/pkg/auth-svc"
	configs "github.com/anazibinurasheed/dmart-api-gateway/pkg/config"
	inventorysvc "github.com/anazibinurasheed/dmart-api-gateway/pkg/inventorysvc"
	util "github.com/anazibinurasheed/dmart-api-gateway/pkg/util"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"  // swagger embed files
	swagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"log"
)

//	@title			 Device Mart Microservice
//	@version		1.0
//	@description	Monolith to microservices
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	Anaz Ibinu Rasheed
//	@contact.url	https://www.linkedin.com/in/anaz-ibinu-rasheed-a2b461253/
//	@contact.email	anazibinurasheed@gmail.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @securitydefinitions.apikey  Bearer
// @in                          header
// @name                Authorization

// @host		localhost:58080
// @BasePath
func main() {
	config, err := configs.LoadConfig()
	if util.HasError(err) {
		log.Fatalln("failed to load config:", err)
	}

	r := gin.New()

	r.Use(gin.Logger())
	r.GET("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler))

	authsvc := authsvc.RegisterRoutes(r, config.AuthSvcPort)
	inventorysvc.RegisterRoutes(r, config.InventorySvcPort, authsvc)

	log.Fatalln(r.Run(config.ApiGatewayPort))
}
