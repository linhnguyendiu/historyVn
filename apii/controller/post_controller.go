package controller

import "github.com/gin-gonic/gin"

type PostController interface {
	Create(ctx *gin.Context)
	GetByTopic(ctx *gin.Context)
	GetByUserId(ctx *gin.Context)
	GetByKeyword(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	LikePost(ctx *gin.Context)
	DisLikePost(ctx *gin.Context)
	GetByPostId(ctx *gin.Context)
	ProcessPosts(ctx *gin.Context)
}
