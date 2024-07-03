package controller

import (
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/model/web"
	"go-pzn-restful-api/service"
	"strconv"

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

func (c *CertificateControllerImpl) GetById(ctx *gin.Context) {
	certId, _ := strconv.Atoi(ctx.Param("certId"))
	userId := ctx.MustGet("current_user").(web.UserResponse).Id
	certResponse := c.CertificateService.FindById(certId, userId)

	ctx.JSON(200,
		helper.APIResponse(200, "Overview course", certResponse))
}

func NewCertificateController(certificateService service.CertificateService) CertificateController {
	return &CertificateControllerImpl{CertificateService: certificateService}
}
