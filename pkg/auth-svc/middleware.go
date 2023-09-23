package authsvc

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/anazibinurasheed/dmart-api-gateway/pkg/auth-svc/pb"
	"github.com/anazibinurasheed/dmart-api-gateway/pkg/authsvc/client"
	util "github.com/anazibinurasheed/dmart-api-gateway/pkg/util"
	"github.com/gin-gonic/gin"
)

type AuthMiddleWare struct {
	svc *client.ServiceClient
}

func InitAuthMiddleware(svc *client.ServiceClient) AuthMiddleWare {
	return InitAuthMiddleware(svc)
}

func (a *AuthMiddleWare) AuthRequired(c *gin.Context) {
	authorization := c.Request.Header.Get("authorization")
	if authorization == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "unauthorized user"})
		return
	}

	token := strings.Split(authorization, "Bearer")
	if len(token) > 2 {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "unauthorized user"})

		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	response, err := a.svc.Client.ValidateToken(
		ctx, &pb.ValidateTokenRequest{
			Token: token[len(token)-1],
		})

	if util.HasError(err) {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if response.Status != http.StatusOK {
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	c.Set("userID", response.UserID)
	c.Next()
}
