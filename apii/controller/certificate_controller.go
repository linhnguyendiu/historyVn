package controller

import "github.com/gin-gonic/gin"

type CertificateController interface {
	Create(ctx *gin.Context)
}
