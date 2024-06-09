package controller

import "github.com/gin-gonic/gin"

type LessonController interface {
	Create(ctx *gin.Context)
	GetByChapterId(ctx *gin.Context)
	Update(ctx *gin.Context)
}
