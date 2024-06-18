package controller

import (
	"fmt"
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/model/web"
	"go-pzn-restful-api/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ImageCourseControllerImpl struct {
	service.ImageCourseService
}

func (c *ImageCourseControllerImpl) UploadImg(ctx *gin.Context) {
	imgId, _ := strconv.Atoi(ctx.Param("imgId"))

	fileHeader, _ := ctx.FormFile("imgAlt")

	pathFile := fmt.Sprintf("assets/images/courses/%d-%s", imgId, fileHeader.Filename)
	uploadImg := c.ImageCourseService.UploadImg(imgId, pathFile)

	ctx.SaveUploadedFile(fileHeader, pathFile)

	ctx.JSON(200,
		helper.APIResponse(200, "Banner is successfully uploaded",
			gin.H{"is_uploaded": uploadImg}),
	)
}

func (c *ImageCourseControllerImpl) GetByCourseId(ctx *gin.Context) {
	courseId, _ := strconv.Atoi(ctx.Param("courseId"))
	imgsResponse := c.ImageCourseService.FindByCourseId(courseId)

	ctx.JSON(200,
		helper.APIResponse(200, "List of chapter titles", imgsResponse))
}

func (c *ImageCourseControllerImpl) Create(ctx *gin.Context) {
	input := web.ImgCourseRequest{}
	err := ctx.ShouldBindJSON(&input)
	helper.PanicIfError(err)

	courseId, _ := strconv.Atoi(ctx.Param("courseId"))
	input.CourseId = courseId

	imgResponse := c.ImageCourseService.Create(input)
	ctx.JSON(200,
		helper.APIResponse(200, "img detail is successfully created", imgResponse))
}

func NewImageCourseController(imageCourseService service.ImageCourseService) ImageCourseController {
	return &ImageCourseControllerImpl{ImageCourseService: imageCourseService}
}
