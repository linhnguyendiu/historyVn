package controller

import (
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/model/web"
	"go-pzn-restful-api/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OptionControllerImpl struct {
	service.OptionService
}

func (c *OptionControllerImpl) GetByQuestionId(ctx *gin.Context) {
	questionId, _ := strconv.Atoi(ctx.Param("questionId"))
	optionsResponse := c.OptionService.FindByQuestionId(questionId)

	ctx.JSON(200,
		helper.APIResponse(200, "List of question titles", optionsResponse))
}

func (c *OptionControllerImpl) Create(ctx *gin.Context) {
	input := web.OptionCreateInput{}
	err := ctx.ShouldBindJSON(&input)
	helper.PanicIfError(err)

	courseId, _ := strconv.Atoi(ctx.Param("courseId"))
	input.CourseId = courseId

	questionId, _ := strconv.Atoi(ctx.Param("questionId"))
	input.QuestionId = questionId

	authorId := ctx.MustGet("current_author").(web.AuthorResponse).Id
	input.AuthorId = authorId

	optionResponse := c.OptionService.Create(input)
	ctx.JSON(200,
		helper.APIResponse(200, "question title is successfully created", optionResponse))
}

func NewOptionController(optionService service.OptionService) OptionController {
	return &OptionControllerImpl{OptionService: optionService}
}
