package authsvc

import (
	"context"
	"fmt"
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
	var role = "suAdmin"
	a.authRequired(c, role)
}

func (a *AuthMiddleware) UserAuth(c *gin.Context) {
	var role = "user"
	a.authRequired(c, role)

}

func (a *AuthMiddleware) authRequired(c *gin.Context, role string) {
	authorization := c.Request.Header.Get("authorization")
	if authorization == "" {
		fmt.Println("1")

		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"StatusCode": 401,
			"msg":        "Unauthorized User",
		})

		return
	}

	token := strings.Split(authorization, "Bearer")
	if len(token) < 2 {

		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"StatusCode": 401,
			"msg":        "Unauthorized User",
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := a.svc.Client.ValidateToken(
		ctx, &pb.ValidateTokenRequest{
			Token: token[1],
			Role:  role,
		})

	if util.HasError(err) {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if resp.Status != http.StatusOK {

		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"StatusCode": 401,
			"data":       resp,
		})
		return
	}

	c.Set("userID", resp.UserID)
	c.Next()
}
