package controller

import (
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/model/web"
	"go-pzn-restful-api/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type LessonControllerImpl struct {
	service.LessonService
}

func (c *LessonControllerImpl) Update(ctx *gin.Context) {
	input := web.LessonCreateInput{}
	err := ctx.ShouldBindJSON(&input)
	helper.PanicIfError(err)
	ltId, _ := strconv.Atoi(ctx.Param("ltId"))

	lessonResponse := c.LessonService.Update(ltId, input)
	ctx.JSON(200,
		helper.APIResponse(200, "Lesson title is successfully updated", lessonResponse))
}

func (c *LessonControllerImpl) GetByChapterId(ctx *gin.Context) {
	chapterId, _ := strconv.Atoi(ctx.Param("chapterId"))
	lessonsResponse := c.LessonService.FindByChapterId(chapterId)

	ctx.JSON(200,
		helper.APIResponse(200, "List of lesson titles", lessonsResponse))
}

func (c *LessonControllerImpl) Create(ctx *gin.Context) {
	input := web.LessonCreateInput{}
	err := ctx.ShouldBindJSON(&input)
	helper.PanicIfError(err)

	chapterId, _ := strconv.Atoi(ctx.Param("chapterId"))
	input.ChapterId = chapterId

	courseId, _ := strconv.Atoi(ctx.Param("courseId"))
	input.CourseId = courseId

	authorId := ctx.MustGet("current_author").(web.AuthorResponse).Id
	input.AuthorId = authorId

	lessonResponse := c.LessonService.Create(input)
	ctx.JSON(200,
		helper.APIResponse(200, "Lesson title is successfully created", lessonResponse))
}

func (c *LessonControllerImpl) UsersCompletedLesson(ctx *gin.Context) {
	user := ctx.MustGet("current_user").(web.UserResponse)
	lessonId, err := strconv.Atoi(ctx.Param("lessonId"))
	helper.PanicIfError(err)

	reponse := c.LessonService.UsersCompletedLesson(user.Id, lessonId)

	ctx.JSON(200,
		helper.APIResponse(200, "Success to enrolled",
			gin.H{"enrolled_by": user.FirstName, "completed_at": reponse.CompletedAt}),
	)
}

func NewLessonController(lessonService service.LessonService) LessonController {
	return &LessonControllerImpl{LessonService: lessonService}
}
