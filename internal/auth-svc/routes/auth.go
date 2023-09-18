package routes

import (
	"github.com/anazibinurasheed/d-api-gateway/internal/auth-svc/payload"
	"github.com/anazibinurasheed/d-api-gateway/internal/auth-svc/pb"
	"github.com/gin-gonic/gin"
)

func CreateAccount(c *gin.Context, req pb.CreateAccountRequest) (pb.CreateAccountResponse, error) {
	var body payload.CreateAccountRequest

}
