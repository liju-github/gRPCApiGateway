package controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/liju-github/EcommerceApiGatewayService/model"
	"github.com/liju-github/EcommerceApiGatewayService/proto/user"
	"google.golang.org/grpc/status"
)

type UserController struct {
	client user.UserServiceClient
}

func NewUserController(client user.UserServiceClient) *UserController {
	return &UserController{client: client}
}

func (uc *UserController) RegisterHandler(c *gin.Context) {
	var req user.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(model.InvalidRequest))
		return
	}

	_, err := uc.client.Register(context.Background(), &req)
	if err != nil {
		st, _ := status.FromError(err)
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(st.Message()))
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(model.RegistrationSuccessful))
}

func (uc *UserController) LoginHandler(c *gin.Context) {
	var req user.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(model.InvalidRequest))
		return
	}

	res, err := uc.client.Login(context.Background(), &req)
	if err != nil {
		st, _ := status.FromError(err)
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(st.Message()))
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": res.Token, "message": model.LoginSuccessful})
}

func (uc *UserController) VerifyEmailHandler(c *gin.Context) {
	var req user.EmailVerificationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(model.InvalidRequest))
		return
	}

	_, err := uc.client.VerifyEmail(context.Background(), &req)
	if err != nil {
		st, _ := status.FromError(err)
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(st.Message()))
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(model.EmailVerificationSuccess))
}

func (uc *UserController) GetProfileHandler(c *gin.Context) {
	var req user.ProfileRequest
	userId := c.GetString("userId")
	req.UserId = userId

	res, err := uc.client.GetProfile(context.Background(), &req)
	if err != nil {
		st, _ := status.FromError(err)
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(st.Message()))
		return
	}

	c.JSON(http.StatusOK, gin.H{"profile": res, "message": model.ProfileRetrieved})
}

func (uc *UserController) UpdateProfileHandler(c *gin.Context) {
	var req user.UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(model.InvalidRequest))
		return
	}

	// Retrieve user ID from the JWT context
	req.UserId = c.GetString("userId")

	res, err := uc.client.UpdateProfile(context.Background(), &req)
	if err != nil {
		st, _ := status.FromError(err)
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(st.Message()))
		return
	}

	c.JSON(http.StatusOK, gin.H{"profile": res.Profile, "message": res.Message})
}
