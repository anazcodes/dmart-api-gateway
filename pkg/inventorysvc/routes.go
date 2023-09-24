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

	router := r

	router.GET("/read-categories", svc.ReadCategories)
	router.GET("/read-products", svc.ReadProducts)

	admin := r.Group("/admin")

	admin.Use(auth.AdminAuth)
	admin.POST("/create-category", svc.CreateCategory)
	admin.POST("/add-product", svc.AddProduct)

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
