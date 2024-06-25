package controller

import (
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/model/web"
	"go-pzn-restful-api/service"

	"github.com/gin-gonic/gin"
)

type CertificateControllerImpl struct {
	service.CertificateService
}

func (c *CertificateControllerImpl) Create(ctx *gin.Context) {
	input := web.CertificateCreateInput{}
	err := ctx.ShouldBindJSON(&input)
	helper.PanicIfError(err)

	certificateResponse := c.CertificateService.Create(input)

	ctx.JSON(200,
		helper.APIResponse(200, "Category has created", certificateResponse))
}

func NewCertificateController(certificateService service.CertificateService) CategoryController {
	return &CertificateControllerImpl{CertificateService: certificateService}
}
