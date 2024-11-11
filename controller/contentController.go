package controller

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/liju-github/EcommerceApiGatewayService/model"
	"github.com/liju-github/EcommerceApiGatewayService/proto/content"
	"google.golang.org/grpc/status"
)

type ContentController struct {
	client content.ContentServiceClient
}

func NewContentController(client content.ContentServiceClient) *ContentController {
	return &ContentController{client: client}
}

func (cc *ContentController) PostQuestionHandler(c *gin.Context) {
	var req content.PostQuestionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(model.InvalidRequest))
		return
	}

	// Set userID from JWT context
	req.UserID = c.GetString("USERID")

	res, err := cc.client.PostQuestion(context.Background(), &req)
	if err != nil {
		st, _ := status.FromError(err)
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(st.Message()))
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": res.Success, "message": res.Message})
}

func (cc *ContentController) GetQuestionsByUserIDHandler(c *gin.Context) {
	userId := c.Query("id")
	if userId == "" {
		c.JSON(http.StatusBadRequest, model.ErrorResponse("Missing question ID"))
		return
	}
	req := &content.GetQuestionsByUserIDRequest{
		UserID: userId,
	}

	res, err := cc.client.GetQuestionsByUserID(context.Background(), req)
	if err != nil {
		st, _ := status.FromError(err)
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(st.Message()))
		return
	}

	c.JSON(http.StatusOK, gin.H{"questions": res.Questions})
}

func (cc *ContentController) GetQuestionsByTagsHandler(c *gin.Context) {
	tagsParam := c.Query("tags")
	if tagsParam == "" {
		c.JSON(http.StatusBadRequest, model.ErrorResponse("tags parameter is required"))
		return
	}

	// Parse JSON array from query parameter
	var tags []string
	if err := json.Unmarshal([]byte(tagsParam), &tags); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse("invalid tags format"))
		return
	}

	req := content.GetQuestionsByTagsRequest{Tags: tags}
	res, err := cc.client.GetQuestionsByTags(context.Background(), &req)
	if err != nil {
		st, _ := status.FromError(err)
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(st.Message()))
		return
	}

	c.JSON(http.StatusOK, gin.H{"questions": res.Questions})
}

func (cc *ContentController) GetQuestionsByWordHandler(c *gin.Context) {
	var req content.GetQuestionsByWordRequest

	searchWord := c.Query("word")
	req.SearchWord = searchWord

	res, err := cc.client.GetQuestionsByWord(context.Background(), &req)
	if err != nil {
		st, _ := status.FromError(err)
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(st.Message()))
		return
	}

	c.JSON(http.StatusOK, gin.H{"questions": res.Questions})
}

func (cc *ContentController) DeleteQuestionHandler(c *gin.Context) {
	var req content.DeleteQuestionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(model.InvalidRequest))
		return
	}

	req.UserID = c.GetString("USERID")

	res, err := cc.client.DeleteQuestion(context.Background(), &req)
	if err != nil {
		st, _ := status.FromError(err)
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(st.Message()))
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": res.Success, "message": res.Message})
}

func (cc *ContentController) GetQuestionByIDHandler(c *gin.Context) {
	questionID := c.Query("id")
	if questionID == "" {
		c.JSON(http.StatusBadRequest, model.ErrorResponse("Missing question ID"))
		return
	}

	req := &content.GetQuestionByIDRequest{
		QuestionID: questionID,
	}

	// Call the gRPC method
	res, err := cc.client.GetQuestionByID(context.Background(), req)
	if err != nil {
		st, _ := status.FromError(err)
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(st.Message()))
		return
	}

	// Respond with the retrieved question
	c.JSON(http.StatusOK, gin.H{"question": res.Question,"answers":res.Answers})
}

func (cc *ContentController) PostAnswerHandler(c *gin.Context) {
	var req content.PostAnswerByQuestionIDRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(model.InvalidRequest))
		return
	}

	// Set userID from JWT context
	req.UserID = c.GetString("USERID")

	res, err := cc.client.PostAnswerByQuestionID(context.Background(), &req)
	if err != nil {
		st, _ := status.FromError(err)
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(st.Message()))
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": res.Success, "message": res.Message})
}

func (cc *ContentController) DeleteAnswerHandler(c *gin.Context) {
	var req content.DeleteAnswerByAnswerIDRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(model.InvalidRequest))
		return
	}

	res, err := cc.client.DeleteAnswerByAnswerID(context.Background(), &req)
	if err != nil {
		st, _ := status.FromError(err)
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(st.Message()))
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": res.Success, "message": res.Message})
}

func (cc *ContentController) UpvoteAnswerHandler(c *gin.Context) {
	var req content.UpvoteAnswerByAnswerIDRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(model.InvalidRequest))
		return
	}
	req.UserID = c.GetString("USERID")

	res, err := cc.client.UpvoteAnswerByAnswerID(context.Background(), &req)
	if err != nil {
		st, _ := status.FromError(err)
		log.Println("upvote error",err.Error())
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(st.Message()))
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": res.Success, "message": res.Message})
}

func (cc *ContentController) DownvoteAnswerHandler(c *gin.Context) {
	var req content.DownvoteAnswerByAnswerIDRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(model.InvalidRequest))
		return
	}

	res, err := cc.client.DownvoteAnswerByAnswerID(context.Background(), &req)
	if err != nil {
		st, _ := status.FromError(err)
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(st.Message()))
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": res.Success, "message": res.Message})
}

func (cc *ContentController) FlagQuestionHandler(c *gin.Context) {
	var req content.FlagQuestionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(model.InvalidRequest))
		return
	}

	// Set userID from JWT context
	req.UserID = c.GetString("USERID")

	res, err := cc.client.FlagQuestion(context.Background(), &req)
	if err != nil {
		st, _ := status.FromError(err)
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(st.Message()))
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": res.Success, "message": res.Message})
}

func (cc *ContentController) FlagAnswerHandler(c *gin.Context) {
	var req content.FlagAnswerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(model.InvalidRequest))
		return
	}

	// Set userID from JWT context
	req.UserID = c.GetString("USERID")
	

	res, err := cc.client.FlagAnswer(context.Background(), &req)
	if err != nil {
		st, _ := status.FromError(err)
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(st.Message()))
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": res.Success, "message": res.Message})
}

func (cc *ContentController) MarkQuestionAsAnsweredHandler(c *gin.Context) {
	var req content.MarkQuestionAsAnsweredRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(model.InvalidRequest))
		return
	}

	res, err := cc.client.MarkQuestionAsAnswered(context.Background(), &req)
	if err != nil {
		st, _ := status.FromError(err)
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(st.Message()))
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": res.Success, "message": res.Message})
}

func (cc *ContentController) GetUserFeedHandler(c *gin.Context) {
	userID := c.GetString("USERID")
	req := &content.GetUserFeedRequest{
		UserID: userID,
	}

	res, err := cc.client.GetUserFeed(context.Background(), req)
	if err != nil {
		st, _ := status.FromError(err)
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(st.Message()))
		return
	}

	c.JSON(http.StatusOK, gin.H{"questions": res.Questions})
}

func (cc *ContentController) AddTagHandler(c *gin.Context) {
	var req content.AddTagRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(model.InvalidRequest))
		return
	}

	res, err := cc.client.AddTag(context.Background(), &req)
	if err != nil {
		st, _ := status.FromError(err)
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(st.Message()))
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": res.Success, "message": res.Message})
}

func (cc *ContentController) RemoveTagHandler(c *gin.Context) {
	var req content.RemoveTagRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(model.InvalidRequest))
		return
	}

	res, err := cc.client.RemoveTag(context.Background(), &req)
	if err != nil {
		st, _ := status.FromError(err)
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(st.Message()))
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": res.Success, "message": res.Message})
}

func (cc *ContentController) SearchHandler(c *gin.Context) {
	var req content.SearchRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(model.InvalidRequest))
		return
	}

	res, err := cc.client.SearchQuestionsAnswersUsers(context.Background(), &req)
	if err != nil {
		st, _ := status.FromError(err)
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(st.Message()))
		return
	}

	c.JSON(http.StatusOK, gin.H{"questions": res.Questions})
}

