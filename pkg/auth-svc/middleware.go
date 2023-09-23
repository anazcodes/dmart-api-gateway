package authsvc

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/anazibinurasheed/dmart-api-gateway/pkg/auth-svc/pb"
	util "github.com/anazibinurasheed/dmart-api-gateway/pkg/util"
	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	svc *ServiceClient
}

func InitAuthMiddleware(svc *ServiceClient) AuthMiddleware {
	return AuthMiddleware{svc: svc}
}

func (a *AuthMiddleware) AdminAuth(c *gin.Context) {
	a.AuthRequired(c)
}

func (a *AuthMiddleware) UserAuth(c *gin.Context) {
	a.AuthRequired(c)

}

func (a *AuthMiddleware) AuthRequired(c *gin.Context) {
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
