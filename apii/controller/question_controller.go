package controller

import "github.com/gin-gonic/gin"

type QuestionController interface {
	Create(ctx *gin.Context)
	GetByCourseId(ctx *gin.Context)
	GetByQuestionId(ctx *gin.Context)
	CreateQuestionWithOptions(ctx *gin.Context)
}
