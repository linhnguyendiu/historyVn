package controller

import "github.com/gin-gonic/gin"

type LessonContentController interface {
	Create(ctx *gin.Context)
	// Update(ctx *gin.Context)
	GetById(ctx *gin.Context)
	GetByLessonId(ctx *gin.Context)
	UploadIllustration(ctx *gin.Context)
}
