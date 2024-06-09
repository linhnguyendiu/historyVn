package controller

import (
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/model/web"
	"go-pzn-restful-api/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ChapterControllerImpl struct {
	service.ChapterService
}

func (c *ChapterControllerImpl) Update(ctx *gin.Context) {
	input := web.ChapterCreateInput{}
	err := ctx.ShouldBindJSON(&input)
	helper.PanicIfError(err)
	ltId, _ := strconv.Atoi(ctx.Param("ltId"))

	chapterResponse := c.ChapterService.Update(ltId, input)
	ctx.JSON(200,
		helper.APIResponse(200, "Chapter title is successfully updated", chapterResponse))
}

func (c *ChapterControllerImpl) GetByCourseId(ctx *gin.Context) {
	courseId, _ := strconv.Atoi(ctx.Param("courseId"))
	chaptersResponse := c.ChapterService.FindByCourseId(courseId)

	ctx.JSON(200,
		helper.APIResponse(200, "List of chapter titles", chaptersResponse))
}

func (c *ChapterControllerImpl) Create(ctx *gin.Context) {
	input := web.ChapterCreateInput{}
	err := ctx.ShouldBindJSON(&input)
	helper.PanicIfError(err)

	courseId, _ := strconv.Atoi(ctx.Param("courseId"))
	input.CourseId = courseId

	authorId := ctx.MustGet("current_author").(web.AuthorResponse).Id
	input.AuthorId = authorId

	chapterResponse := c.ChapterService.Create(input)
	ctx.JSON(200,
		helper.APIResponse(200, "chapter title is successfully created", chapterResponse))
}

func NewChapterController(chapterService service.ChapterService) ChapterController {
	return &ChapterControllerImpl{ChapterService: chapterService}
}
