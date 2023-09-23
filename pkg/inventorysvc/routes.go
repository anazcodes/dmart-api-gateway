package inventorysvc

import (
	authsvc "github.com/anazibinurasheed/d-api-gateway/internal/auth-svc"
	"github.com/anazibinurasheed/d-api-gateway/internal/inventorysvc/client"
	"github.com/anazibinurasheed/d-api-gateway/internal/inventorysvc/handler"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, inventorySvcPort string, auth *authsvc.Ha) {
	svc := &handler.Client{
		ServiceClient: client.ServiceClient{
			Client: client.InitServiceClient(inventorySvcPort),
		},
	}

	router := r.Group("")
	router.POST("/create-category", svc.CreateCategory)
	router.POST("/read-categories", svc.ReadCategories)
	router.POST("/add-product", svc.AddProduct)
	router.GET("/read-products", svc.ReadProducts)

}
