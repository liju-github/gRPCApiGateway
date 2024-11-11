package controller

// import (
// 	"context"
// 	"log"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"github.com/liju-github/EcommerceApiGatewayService/model"
// 	"github.com/liju-github/EcommerceApiGatewayService/proto/admin"
// 	"github.com/liju-github/EcommerceApiGatewayService/proto/content"
// 	"github.com/liju-github/EcommerceApiGatewayService/proto/user"
// 	"google.golang.org/grpc/status"
// )

// type AdminController struct {
// 	client admin
// }

// func NewAdminController(client admin.AdminServiceClient) *AdminController {
// 	return &AdminController{client: client}
// }

// func (uc *AdminController) BanAdminHandler(c *gin.Context) {
// 	var req admin.BanAdminRequest
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, model.ErrorResponse(model.InvalidRequest))
// 		return
// 	}
// 	if req.AdminId == "" {
// 		c.JSON(http.StatusBadRequest, &admin.BanUserResponse{
// 			Success: false,
// 			Message: "adminID is not present",
// 		})
// 		return
// 	}

// 	res, err := uc.client.BanAdmin(c, &req)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, res)
// 		return
// 	}

// 	c.JSON(http.StatusOK, res)
// }

// func (uc *AdminController) UnBanAdminHandler(c *gin.Context) {
// 	var req user.UnBanUserRequest
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, model.ErrorResponse(model.InvalidRequest))
// 		return
// 	}
// 	if req.UserId == "" {
// 		c.JSON(http.StatusBadRequest, &user.UnBanUserResponse{
// 			Success: false,
// 			Message: "userID is not present",
// 		})
// 		return
// 	}

// 	res, err := uc.client.UnBanAdmin(c, &req)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, res)
// 		return
// 	}

// 	c.JSON(http.StatusOK, res)
// }

// func (cc *AdminController) GetAllFlaggedAnswers(c *gin.Context) {
// 	var req content.GetFlaggedAnswersRequest

// 	res, err := cc.client.GetFlaggedAnswers(context.Background(), &req)
// 	if err != nil {
// 		st, _ := status.FromError(err)
// 		c.JSON(http.StatusInternalServerError, model.ErrorResponse(st.Message()))
// 		return
// 	}
// 	log.Println("printing", res.FlaggedAnswers)

// 	c.JSON(http.StatusOK, gin.H{
// 		"data": res,
// 	})

// }

// func (cc *AdminController) GetAllFlaggedQuestions(c *gin.Context) {
// 	var req content.GetFlaggedQuestionsRequest

// 	res, err := cc.client.GetFlaggedQuestions(context.Background(), &req)
// 	if err != nil {
// 		st, _ := status.FromError(err)
// 		c.JSON(http.StatusInternalServerError, model.ErrorResponse(st.Message()))
// 		return
// 	}

// 	log.Println("printing", res)

// 	c.JSON(http.StatusOK, gin.H{
// 		"data": res,
// 	})

// }
