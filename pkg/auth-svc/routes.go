package authsvc

import (
	"github.com/anazibinurasheed/dmart-api-gateway/pkg/auth-svc/handler"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, authSvcAddress string) *ServiceClient {

	svc := &ServiceClient{
		Client: InitServiceClient(authSvcAddress),
	}

	routes := r.Group("/auth")
	routes.POST("/create-account", svc.CreateAccount)
	routes.POST("/login", svc.UserLogin)
	routes.POST("/login/admin", svc.AdminLogin)
	return svc
}

func (ac *ServiceClient) CreateAccount(c *gin.Context) {
	handler.CreateAccount(c, ac.Client)
}

func (ac *ServiceClient) UserLogin(c *gin.Context) {
	handler.UserLogin(c, ac.Client)
}

func (ac *ServiceClient) AdminLogin(c *gin.Context) {
	handler.AdminLogin(c, ac.Client)
}
