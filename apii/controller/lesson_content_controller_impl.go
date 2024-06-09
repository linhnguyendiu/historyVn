package controller

import (
	"fmt"
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/model/web"
	"go-pzn-restful-api/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type LessonContentControllerImpl struct {
	service.LessonContentService
}

func (c *LessonContentControllerImpl) GetById(ctx *gin.Context) {
	isUserHas := ctx.MustGet("isUserHas").(bool)

	if isUserHas == false {
		ctx.AbortWithStatusJSON(200,
			helper.APIResponse(
				200, "List of lesson contents",
				gin.H{"is_user_has": isUserHas, "message": "You must unlock this course first"},
			),
		)
		return
	}

	lcId, _ := strconv.Atoi(ctx.Param("lcId"))
	findById := c.LessonContentService.FindById(lcId)

	ctx.JSON(200, helper.APIResponse(200, "Detail of lesson content", findById))
}

func (c *LessonContentControllerImpl) GetByLessonId(ctx *gin.Context) {
	ltId, err := strconv.Atoi(ctx.Param("ltId"))
	helper.PanicIfError(err)

	lessonContentsResponse := c.LessonContentService.FindByLessonId(ltId)

	ctx.JSON(200,
		helper.APIResponse(200, "List of lesson contents", lessonContentsResponse),
	)
}

// func (c *LessonContentControllerImpl) Update(ctx *gin.Context) {
// 	input := web.LessonContentUpdateInput{}
// 	err := ctx.ShouldBindJSON(&input)
// 	authorId := ctx.MustGet("current_author").(web.AuthorResponse).Id
// 	courseId, err := strconv.Atoi(ctx.Param("courseId"))
// 	lcId, err := strconv.Atoi(ctx.Param("lcId"))
// 	helper.PanicIfError(err)

// 	input.AuthorId = authorId
// 	input.CourseId = courseId

// 	fileHeader, err := ctx.FormFile("content")
// 	if err != nil {
// 		input.Content = ""
// 		lessonContentResponse := c.LessonContentService.Update(lcId, input)
// 		ctx.JSON(200,
// 			helper.APIResponse(200, "Lesson content successfully updated", lessonContentResponse),
// 		)
// 		return
// 	}
// 	pathContent := fmt.Sprintf("assets/contents/%s", fileHeader.Filename)
// 	input.Content = pathContent
// 	err = ctx.SaveUploadedFile(fileHeader, pathContent)
// 	helper.PanicIfError(err)

// 	lessonContentResponse := c.LessonContentService.Update(lcId, input)

// 	ctx.JSON(200,
// 		helper.APIResponse(200, "Lesson content successfully updated", lessonContentResponse),
// 	)

// }

func (c *LessonContentControllerImpl) Create(ctx *gin.Context) {
	input := web.LessonContentCreateInput{}
	err := ctx.ShouldBindJSON(&input)
	helper.PanicIfError(err)

	authorId := ctx.MustGet("current_author").(web.AuthorResponse).Id
	courseId, _ := strconv.Atoi(ctx.Param("courseId"))
	ltId, _ := strconv.Atoi(ctx.Param("ltId"))

	input.AuthorId = authorId
	input.CourseId = courseId
	input.LessonId = ltId

	lessonContentResponse := c.LessonContentService.Create(input)

	ctx.JSON(200,
		helper.APIResponse(200, "Lesson content successfully created", lessonContentResponse),
	)
}

func (c *LessonContentControllerImpl) UploadIllustration(ctx *gin.Context) {
	lcId, _ := strconv.Atoi(ctx.Param("lcId"))

	fileHeader, _ := ctx.FormFile("illustration")

	pathFile := fmt.Sprintf("assets/images/illustration/%d-%s", lcId, fileHeader.Filename)
	uploadIllustration := c.LessonContentService.UploadIllustration(lcId, pathFile)

	ctx.SaveUploadedFile(fileHeader, pathFile)

	ctx.JSON(200,
		helper.APIResponse(200, "Illustration is successfully uploaded",
			gin.H{"is_uploaded": uploadIllustration}),
	)
}

func NewLessonContentController(lessonContentService service.LessonContentService) LessonContentController {
	return &LessonContentControllerImpl{LessonContentService: lessonContentService}
}
