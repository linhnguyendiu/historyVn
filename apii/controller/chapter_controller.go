package controller

import "github.com/gin-gonic/gin"

type ChapterController interface {
	Create(ctx *gin.Context)
	GetByCourseId(ctx *gin.Context)
	Update(ctx *gin.Context)
}
