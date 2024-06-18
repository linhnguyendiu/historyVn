package controller

import "github.com/gin-gonic/gin"

type ImageCourseController interface {
	Create(ctx *gin.Context)
	GetByCourseId(ctx *gin.Context)
	UploadImg(ctx *gin.Context)
}
