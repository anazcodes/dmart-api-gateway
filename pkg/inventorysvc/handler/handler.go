package handler

import (
	"context"
	"fmt"
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

	if err := payload.ValidateStruct(&body); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

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
	fmt.Println("1")

	if !util.BindRequest(c, &body) {
		return
	}

	if err := payload.ValidateStruct(&body); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	fmt.Println("2")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	response, err := i.AddProduct(ctx, &pb.AddProductRequest{
		CategoryID:  body.CategoryID,
		Name:        body.Name,
		Description: body.Description,
		Image:       body.Image,
		Price:       body.Price,
	})
	fmt.Println("3")

	if util.HasError(err) {
		c.JSON(http.StatusBadGateway, err.Error())
		return
	}
	fmt.Println("4")

	c.JSON(http.StatusCreated, response)
	fmt.Println("5")

}

func ReadProducts(c *gin.Context, i pb.InventoryServiceClient) {

	page, count, err := util.GetPageNCount(c)
	if util.HasError(err) {
		response := util.Response(http.StatusBadRequest, "Invalid parameters provided for pagination", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	response, err := i.ReadProducts(ctx, &pb.Request{
		Page:  int64(page),
		Count: int64(count),
	})

	if util.HasError(err) {
		c.JSON(http.StatusBadGateway, err.Error())
		return
	}

	c.JSON(http.StatusOK, response)
}
