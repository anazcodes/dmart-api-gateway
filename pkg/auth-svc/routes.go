package authsvc

import (
	"github.com/anazibinurasheed/dmart-api-gateway/pkg/auth-svc/client"
	"github.com/anazibinurasheed/dmart-api-gateway/pkg/auth-svc/handler"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, authSvcAddress string) *handler.Client {

	svc := &handler.Client{
		ServiceClient: client.ServiceClient{
			Client: client.InitServiceClient(authSvcAddress),
		},
	}

	routes := r.Group("/auth")
	routes.POST("/create-account", svc.CreateAccount)
	routes.POST("/login", svc.UserLogin)
	routes.POST("/login/admin", svc.AdminLogin)
	return svc
}
