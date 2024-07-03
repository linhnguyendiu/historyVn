package controller

import "github.com/gin-gonic/gin"

type RewardController interface {
	Create(ctx *gin.Context)
	GetByUserId(ctx *gin.Context)
}
