package handler

import (
	"github.com/anazibinurasheed/d-api-gateway/internal/auth-svc/client"
	"github.com/gin-gonic/gin"
)
type Client struct {
	client.ServiceClient
}

func (ac *Client) CreateAccount(c *gin.Context) {
	CreateAccount(c, ac.Client)
}

func (ac *Client) UserLogin(c *gin.Context) {
	UserLogin(c, ac.Client)
}

func (ac *Client) AdminLogin(c *gin.Context) {
	AdminLogin(c, ac.Client)
}
