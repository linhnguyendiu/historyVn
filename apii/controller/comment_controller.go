package controller

import "github.com/gin-gonic/gin"

type CommentController interface {
	Create(ctx *gin.Context)
	GetCommentsByPostId(ctx *gin.Context)
	LikeComment(ctx *gin.Context)
	DisLikeComment(ctx *gin.Context)
	GetByCommentFatherId(ctx *gin.Context)
	GetCommentsByUserId(ctx *gin.Context)
	ProcessComments(ctx *gin.Context)
}
