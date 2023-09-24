package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/anazibinurasheed/dmart-api-gateway/pkg/auth-svc/payload"
	"github.com/anazibinurasheed/dmart-api-gateway/pkg/auth-svc/pb"
	util "github.com/anazibinurasheed/dmart-api-gateway/pkg/util"
	"github.com/gin-gonic/gin"
)

// UserLogin godoc.
//
//	@Summary		Create account
//	@Description	User can create an account using this endpoint
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			body	body		payload.CreateAccountRequest	true	"login credentials"
//	@Success		201		{object}	pb.CreateAccountResponse
//	@Failure		502		{object}	util.response
//	@Failure		400		{object}	util.response
//	@Failure		417		{object}	util.response
//	@Router			/auth/create-account [post]
func CreateAccount(c *gin.Context, asc pb.AuthServiceClient) {
	var body payload.CreateAccountRequest

	if !util.BindRequest(c, &body) {
		return
	}

	if !util.ValidateStruct(c, &body) {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := asc.CreateAccount(ctx, &pb.CreateAccountRequest{
		Username: body.Username,
		Email:    body.Email,
		Phone:    body.Phone,
		Password: body.Password,
	})

	if util.RpcHasError(c, err) {
		return
	}

	c.JSON(int(res.Status), res)
}

// UserLogin godoc.
//
//	@Summary		User login
//	@Description	User can login using this endpoint
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			body	body		payload.UserLoginRequest	true	"login credentials"
//	@Success		200		{object}	pb.UserLoginResponse
//	@Failure		400		{object}	util.response
//	@Failure		417		{object}	util.response
//	@Failure		502		{object}	util.response
//	@Router			/auth/login [post]
func UserLogin(c *gin.Context, asc pb.AuthServiceClient) {
	var body payload.UserLoginRequest

	if !util.BindRequest(c, &body) {
		return
	}
	if !util.ValidateStruct(c, &body) {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := asc.UserLogin(ctx, &pb.UserLoginRequest{
		LoginInput: body.LoginInput,
		Password:   body.Password,
	})

	if util.RpcHasError(c, err) {
		return
	}

	c.Request.Header.Set("authorization", util.BuildBearerToken(res.Token))

	c.JSON(int(res.Status), res)
}

// AdminLogin godoc.
//
//	@Summary		Admin login
//	@Description    Admin can login using this endpoint
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			body	body		payload.AdminLoginRequest	true	"login credentials"
//	@Success		200		{object}	pb.AdminLoginResponse
//	@Failure		401		{object}	util.response
//	@Failure		400		{object}	util.response
//	@Failure		500		{object}	util.response
//	@Router			/auth/login/admin [post]
func AdminLogin(c *gin.Context, asc pb.AuthServiceClient) {
	var body payload.AdminLoginRequest

	if !util.BindRequest(c, &body) {
		return
	}
	if !util.ValidateStruct(c, &body) {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := asc.AdminLogin(ctx, &pb.AdminLoginRequest{
		Username: body.Username,
		Password: body.Password,
	})

	if util.RpcHasError(c, err) {
		return
	}

	c.Request.Header.Set("authorization", util.BuildBearerToken(res.Token))

	c.JSON(int(res.Status), res)
}

// Logout godoc.
//
//	@Summary		Logout
//	@Description    Logout endpoint
//	@Tags			auth
//	@Produce		json
//	@Success		202		{object}	util.response
//	@Router			/auth/logout [post]
func Logout(c *gin.Context) {
	c.Request.Header.Set("authorization", "")
	response := util.Response(http.StatusAccepted, "success, logged out", nil, nil)
	c.JSON(http.StatusAccepted, response)
}
