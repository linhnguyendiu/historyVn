package controller

import (
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/model/web"
	"go-pzn-restful-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentControllerImpl struct {
	service.CommentService
}

func (c *CommentControllerImpl) GetCommentsByPostId(ctx *gin.Context) {
	param := ctx.Param("postId")
	postId, _ := strconv.Atoi(param)
	commentResponse := c.CommentService.FindByPostId(postId)
	ctx.JSON(200,
		helper.APIResponse(200, "List of comments", commentResponse),
	)
}

func (c *CommentControllerImpl) GetCommentsByUserId(ctx *gin.Context) {
	userId := ctx.MustGet("current_user").(web.UserResponse).Id
	commentResponse := c.CommentService.FindByUserId(userId)
	ctx.JSON(200,
		helper.APIResponse(200, "List of comments", commentResponse),
	)
}

func (c *CommentControllerImpl) GetByCommentFatherId(ctx *gin.Context) {
	commentId, _ := strconv.Atoi(ctx.Param("commentId"))
	commentResponse := c.CommentService.FindByComId(commentId)
	ctx.JSON(200,
		helper.APIResponse(200, "post detail", commentResponse),
	)
}

func (c *CommentControllerImpl) Create(ctx *gin.Context) {
	request := web.CommentCreateInput{}
	err := ctx.ShouldBindJSON(&request)
	helper.PanicIfError(err)

	userId := ctx.MustGet("current_user").(web.UserResponse).Id
	request.UserId = userId

	commentResponse := c.CommentService.Create(request)
	ctx.JSON(200,
		helper.APIResponse(200, "Comment has been created", commentResponse),
	)
}

func (c *CommentControllerImpl) LikeComment(ctx *gin.Context) {
	commentId, _ := strconv.Atoi(ctx.Param("commentId"))

	userId := ctx.MustGet("current_user").(web.UserResponse).Id

	likes, didUserLike, err := c.CommentService.LikeComment(ctx, userId, commentId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":      "Comment liked successfully",
		"likes":       likes,
		"didUserLike": didUserLike,
	})
}

func (c *CommentControllerImpl) DisLikeComment(ctx *gin.Context) {
	commentId, _ := strconv.Atoi(ctx.Param("commentId"))

	userId := ctx.MustGet("current_user").(web.UserResponse).Id

	disLikes, didUserDisLike, err := c.CommentService.DisLikeComment(ctx, userId, commentId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":      "Comment disliked successfully",
		"dislikes":    disLikes,
		"didUserLike": didUserDisLike,
	})
}

func (c *CommentControllerImpl) ProcessComments(ctx *gin.Context) {
	reponses := c.CommentService.ProcessComments()
	ctx.JSON(200,
		helper.APIResponse(200, "Calculate point all post successfull", reponses),
	)
}

func NewCommentController(commentService service.CommentService) CommentController {
	return &CommentControllerImpl{CommentService: commentService}
}
