package handler

import (
	"github.com/anazibinurasheed/d-api-gateway/internal/client"
	"github.com/gin-gonic/gin"
)

type Client struct {
	client.ServiceClient
}

func (ac *Client) CreateAccount(c *gin.Context) {
	CreateAccount(c, ac.Client)
}
