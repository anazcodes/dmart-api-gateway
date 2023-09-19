package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/anazibinurasheed/d-api-gateway/internal/auth-svc/payload"
	"github.com/anazibinurasheed/d-api-gateway/internal/auth-svc/pb"
	util "github.com/anazibinurasheed/d-api-gateway/internal/utils"
	"github.com/gin-gonic/gin"
)

func CreateAccount(c *gin.Context, asc pb.AuthServiceClient) {
	var body payload.CreateAccountRequest

	util.BindRequest(c, &body)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := asc.CreateAccount(ctx, &pb.CreateAccountRequest{
		Username:  body.Username,
		Email:     body.Email,
		Phone:     body.Phone,
		Password:  body.Password,
		Password2: body.Password,
	})

	if err != nil {
		util.LogInput(err.Error())
		c.JSON(http.StatusBadGateway, err)
		return
	}

	c.JSON(int(res.Status), res)
}

// showing:

// rpc error: code = Canceled desc = grpc: the client connection is closing
