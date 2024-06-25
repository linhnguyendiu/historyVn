package controller

import (
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/model/web"
	"go-pzn-restful-api/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type QuestionControllerImpl struct {
	service.QuestionService
}

func (c *QuestionControllerImpl) GetByCourseId(ctx *gin.Context) {
	courseId, _ := strconv.Atoi(ctx.Param("courseId"))
	questionsResponse := c.QuestionService.FindByCourseId(courseId)

	ctx.JSON(200,
		helper.APIResponse(200, "List of question by courseId", questionsResponse))
}

func (c *QuestionControllerImpl) GetByQuestionId(ctx *gin.Context) {
	questionId, _ := strconv.Atoi(ctx.Param("questionId"))
	questionResponse := c.QuestionService.FindById(questionId)

	ctx.JSON(200,
		helper.APIResponse(200, "List of question", questionResponse))
}

func (c *QuestionControllerImpl) Create(ctx *gin.Context) {
	input := web.QuestionCreateInput{}
	err := ctx.ShouldBindJSON(&input)
	helper.PanicIfError(err)

	courseId, _ := strconv.Atoi(ctx.Param("courseId"))
	input.CourseId = courseId

	authorId := ctx.MustGet("current_author").(web.AuthorResponse).Id
	input.AuthorId = authorId

	questionResponse := c.QuestionService.Create(input)
	ctx.JSON(200,
		helper.APIResponse(200, "question title is successfully created", questionResponse))
}

func (c *QuestionControllerImpl) CreateQuestionWithOptions(ctx *gin.Context) {
	input := web.ListQuestionCreateInput{}
	err := ctx.ShouldBindJSON(&input)
	helper.PanicIfError(err)

	courseId, _ := strconv.Atoi(ctx.Param("courseId"))
	input.CourseId = courseId

	authorId := ctx.MustGet("current_author").(web.AuthorResponse).Id
	input.AuthorId = authorId

	response := c.QuestionService.CreateWithOptions(input)
	ctx.JSON(200,
		helper.APIResponse(200, "question title is successfully created", response))
}

func NewQuestionController(questionService service.QuestionService) QuestionController {
	return &QuestionControllerImpl{QuestionService: questionService}
}
