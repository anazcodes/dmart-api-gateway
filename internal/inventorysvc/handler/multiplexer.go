package handler

import (
	"github.com/anazibinurasheed/d-api-gateway/internal/inventorysvc/client"
	"github.com/gin-gonic/gin"
)

type Client struct {
	client.ServiceClient
}

func (inv *Client) CreateCategory(c *gin.Context) {
	CreateCategory(c, inv.Client)
}

func (inv *Client) ReadCategories(c *gin.Context) {
	ReadCategories(c, inv.Client)
}
