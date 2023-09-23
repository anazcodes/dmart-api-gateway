package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/anazibinurasheed/d-api-gateway/internal/auth-svc/payload"
	"github.com/anazibinurasheed/d-api-gateway/internal/auth-svc/pb"
	util "github.com/anazibinurasheed/d-api-gateway/internal/util"
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

	if util.HasError(err) {
		util.Logger(err.Error())
		c.JSON(http.StatusBadGateway, err.Error())
		return
	}

	c.JSON(int(res.Status), res)
}

func UserLogin(c *gin.Context, asc pb.AuthServiceClient) {
	var body payload.UserLoginRequest

	util.BindRequest(c, &body)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := asc.UserLogin(ctx, &pb.UserLoginRequest{
		LoginInput: body.LoginInput,
		Password:   body.Password,
	})

	if err != nil {
		util.Logger(err.Error())
		c.JSON(http.StatusBadGateway, err)
		return
	}

	c.JSON(int(res.Status), res)
}

func AdminLogin(c *gin.Context, asc pb.AuthServiceClient) {
	var body payload.AdminLoginRequest

	util.BindRequest(c, &body)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := asc.AdminLogin(ctx, &pb.AdminLoginRequest{
		Username: body.Username,
		Password: body.Password,
	})

	if err != nil {
		util.Logger(err.Error())
		c.JSON(http.StatusBadGateway, err)
		return
	}

	c.JSON(int(res.Status), res)
}
