package inventorysvc

import (
	authsvc "github.com/anazibinurasheed/dmart-api-gateway/pkg/auth-svc"
	"github.com/anazibinurasheed/dmart-api-gateway/pkg/inventorysvc/handler"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, inventorySvcPort string, authSvc *authsvc.ServiceClient) {
	auth := authsvc.InitAuthMiddleware(authSvc)
	svc := &ServiceClient{
		Client: InitServiceClient(inventorySvcPort),
	}

	router := r.Group("")
	router.Use(auth.AuthRequired)
	router.POST("/create-category", svc.CreateCategory)
	router.POST("/read-categories", svc.ReadCategories)
	router.POST("/add-product", svc.AddProduct)
	router.GET("/read-products", svc.ReadProducts)

}

func (inv *ServiceClient) CreateCategory(c *gin.Context) {
	handler.CreateCategory(c, inv.Client)
}

func (inv *ServiceClient) ReadCategories(c *gin.Context) {
	handler.ReadCategories(c, inv.Client)
}

func (inv *ServiceClient) AddProduct(c *gin.Context) {
	handler.AddProduct(c, inv.Client)
}

func (inv *ServiceClient) ReadProducts(c *gin.Context) {
	handler.ReadProducts(c, inv.Client)
}
