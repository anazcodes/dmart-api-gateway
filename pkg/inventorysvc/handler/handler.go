package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/anazibinurasheed/dmart-api-gateway/pkg/inventorysvc/payload"
	"github.com/anazibinurasheed/dmart-api-gateway/pkg/inventorysvc/pb"
	"github.com/anazibinurasheed/dmart-api-gateway/pkg/util"
	"github.com/gin-gonic/gin"
)

// @Summary Create category
// @Description Admin can create category
// @Tags inventory-admin
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param			body		body	payload.CreateCategoryRequest	true	"Category object"
// @Success		200			{object}	pb.AddProductResponse
// @Failure 	400    {object} util.response
// @Failure 	417    {object} util.response
// @Failure 	502    {object} util.response
// @Router			/admin/create-category [post]
func CreateCategory(c *gin.Context, inv pb.InventoryServiceClient) {
	var body payload.CreateCategoryRequest
	if !util.BindRequest(c, &body) {
		return
	}
	if !util.ValidateStruct(c, &body) {
		return
	}

	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	res, err := inv.CreateCategory(ctx, &pb.CreateCategoryRequest{
		Name: body.Name,
	})

	if util.RpcHasError(c, err) {
		return
	}

	c.JSON(int(res.Status), res)
}

// @Summary View all categories
// @Description User can view all available categories
// @Tags inventory
// @Produce json
// @Param			page	query	int	true	"Page number"
// @Param			count	query	int	true	"Number of items per page"
// @Success		200			{object}	pb.ReadCategoriesResponse
// @Failure 	400    {object} util.response
// @Failure 	417    {object} util.response
// @Failure 	502    {object} util.response
// @Router			/read-categories [get]
func ReadCategories(c *gin.Context, i pb.InventoryServiceClient) {
	page, count, err := util.GetPageNCount(c)
	if util.ErrorInPageInfo(c, err) {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	data, err := i.ReadCategories(ctx, &pb.Request{
		Page:  int64(page),
		Count: int64(count),
	})

	if util.RpcHasError(c, err) {
		return
	}

	c.JSON(http.StatusOK, data)
}

// @Summary Add product
// @Description Admin can add product
// @Tags inventory-admin
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param			body		body	payload.CreateProductRequest	true	"Product object"
// @Success		200			{object}	pb.AddProductResponse
// @Failure 	400    {object} util.response
// @Failure 	417    {object} util.response
// @Failure 	502    {object} util.response
// @Router			/admin/add-product [post]
func AddProduct(c *gin.Context, i pb.InventoryServiceClient) {
	var body payload.CreateProductRequest

	if !util.BindRequest(c, &body) {
		return
	}
	if !util.ValidateStruct(c, &body) {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := i.AddProduct(ctx, &pb.AddProductRequest{
		CategoryID:  body.CategoryID,
		Name:        body.Name,
		Description: body.Description,
		Price:       body.Price,
	})

	if util.RpcHasError(c, err) {
		return
	}

	c.JSON(http.StatusCreated, res)
}

// @Summary View all products
// @Description User can view all available products
// @Tags inventory
// @Produce json
// @Param			page	query	int	true	"Page number"
// @Param			count	query	int	true	"Number of items per page"
// @Success		200			{object}	pb.ReadProductsResponse
// @Failure 	400    {object} util.response
// @Failure 	417    {object} util.response
// @Failure 	502    {object} util.response
// @Router			/read-products [get]
func ReadProducts(c *gin.Context, i pb.InventoryServiceClient) {
	page, count, err := util.GetPageNCount(c)
	if util.ErrorInPageInfo(c, err) {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	response, err := i.ReadProducts(ctx, &pb.Request{
		Page:  int64(page),
		Count: int64(count),
	})

	if util.RpcHasError(c, err) {
		return
	}

	c.JSON(http.StatusOK, response)
}
