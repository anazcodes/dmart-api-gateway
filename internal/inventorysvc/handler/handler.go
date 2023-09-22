package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/anazibinurasheed/d-api-gateway/internal/inventorysvc/payload"
	"github.com/anazibinurasheed/d-api-gateway/internal/inventorysvc/pb"
	"github.com/anazibinurasheed/d-api-gateway/internal/util"
	"github.com/gin-gonic/gin"
)

func CreateCategory(c *gin.Context, inv pb.InventoryServiceClient) {
	var body payload.CreateCategoryRequest
	util.BindRequest(c, &body)

	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	res, err := inv.CreateCategory(ctx, &pb.CreateCategoryRequest{
		Name: body.Name,
	})

	if util.HasError(err) {
		c.JSON(http.StatusBadGateway, err.Error())
		return
	}
	c.JSON(int(res.Status), res)
}

func ReadCategories(c *gin.Context, i pb.InventoryServiceClient) {

	page, count, err := util.GetPageNCount(c)
	if util.HasError(err) {
		response := util.Response(http.StatusBadRequest, "Invalid parameters provided for pagination", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	data, err := i.ReadCategories(ctx, &pb.Request{
		Page:  int64(page),
		Count: int64(count),
	})

	if util.HasError(err) {
		c.JSON(http.StatusBadGateway, err.Error())
		return
	}

	c.JSON(http.StatusOK, data)
}

func AddProduct(c *gin.Context, i pb.InventoryServiceClient) {
	var body payload.CreateProductRequest
	util.BindRequest(c, &body)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	i.AddProduct(ctx, &pb.AddProductRequest{})

}

func ReadProducts(c *gin.Context, i pb.InventoryServiceClient) {

}
