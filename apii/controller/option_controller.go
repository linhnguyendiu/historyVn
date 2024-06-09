package controller

import "github.com/gin-gonic/gin"

type OptionController interface {
	Create(ctx *gin.Context)
	GetByQuestionId(ctx *gin.Context)
}
