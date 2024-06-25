package controller

import (
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/model/web"
	"go-pzn-restful-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CourseControllerImpl struct {
	service.CourseService
}

func (c *CourseControllerImpl) GetByCategory(ctx *gin.Context) {
	courseResponses := c.CourseService.FindByCategory(ctx.Param("categoryName"))

	ctx.JSON(200,
		helper.APIResponse(200, "List of courses", courseResponses),
	)
}

func (c *CourseControllerImpl) GetByTypeAndCategory(ctx *gin.Context) {
	courseResponses := c.CourseService.FindByTypeAndCategory(ctx.Param("type"), ctx.Param("cateName"))

	ctx.JSON(200,
		helper.APIResponse(200, "List of courses", courseResponses),
	)
}

func (c *CourseControllerImpl) GetByUserId(ctx *gin.Context) {
	userId := ctx.MustGet("current_user").(web.UserResponse).Id
	courseResponses := c.CourseService.FindAllCourseIdByUserId(userId)

	ctx.JSON(200,
		helper.APIResponse(200, "List of courses", courseResponses),
	)
}

func (c *CourseControllerImpl) GetResultByUserId(ctx *gin.Context) {
	userId := ctx.MustGet("current_user").(web.UserResponse).Id
	courseId, err := strconv.Atoi(ctx.Param("courseId"))
	helper.PanicIfError(err)
	courseResponses := c.CourseService.FindResultById(userId, courseId)

	ctx.JSON(200,
		helper.APIResponse(200, "Result", courseResponses),
	)
}

func (c *CourseControllerImpl) UserEnrolled(ctx *gin.Context) {
	user := ctx.MustGet("current_user").(web.UserResponse)
	courseId, err := strconv.Atoi(ctx.Param("courseId"))
	helper.PanicIfError(err)

	c.CourseService.UserEnrolled(user.Id, courseId)

	ctx.JSON(200,
		helper.APIResponse(200, "Success to enrolled",
			gin.H{"enrolled_by": user.FirstName}),
	)
}

func (c *CourseControllerImpl) GetAll(ctx *gin.Context) {
	courseResponses := c.CourseService.FindAll()
	ctx.JSON(200,
		helper.APIResponse(200, "List of courses", courseResponses),
	)
}

func (c *CourseControllerImpl) GetByKeyword(ctx *gin.Context) {
	courseResponses, err := c.CourseService.FindByKeyword(ctx.Param("keywords"))
	helper.PanicIfError(err)
	ctx.JSON(http.StatusOK,
		helper.APIResponse(200, "List of posts", courseResponses),
	)
}

func (c *CourseControllerImpl) GetTop3Course(ctx *gin.Context) {
	courseResponses := c.CourseService.FindTop3Coures()
	ctx.JSON(http.StatusOK,
		helper.APIResponse(200, "List of top courses", courseResponses),
	)
}

func (c *CourseControllerImpl) GetByAuthorId(ctx *gin.Context) {
	param := ctx.Param("authorId")
	authorId, _ := strconv.Atoi(param)
	courseResponse := c.CourseService.FindByAuthorId(authorId)
	ctx.JSON(200,
		helper.APIResponse(200, "List of courses", courseResponse),
	)
}

func (c *CourseControllerImpl) GetByType(ctx *gin.Context) {
	courseResponse := c.CourseService.FindByType(ctx.Param("type"))
	ctx.JSON(http.StatusOK,
		helper.APIResponse(200, "Course detail", courseResponse),
	)
}

func (c *CourseControllerImpl) Create(ctx *gin.Context) {
	request := web.CourseCreateInput{}
	err := ctx.ShouldBindJSON(&request)
	helper.PanicIfError(err)

	authorId := ctx.MustGet("current_author").(web.AuthorResponse).Id
	request.AuthorId = authorId

	courseResponse := c.CourseService.Create(request)
	ctx.JSON(200,
		helper.APIResponse(200, "Course has been created", courseResponse),
	)
}

func (c *CourseControllerImpl) GetExamScore(ctx *gin.Context) {
	request := web.ExamRequest{}
	err := ctx.ShouldBindJSON(&request)
	helper.PanicIfError(err)

	courseId, err := strconv.Atoi(ctx.Param("courseId"))
	helper.PanicIfError(err)
	request.CourseId = courseId

	userId := ctx.MustGet("current_user").(web.UserResponse).Id
	request.UserId = userId

	examResultResponse := c.CourseService.GetScore(ctx, request)

	ctx.JSON(http.StatusOK, gin.H{"result": examResultResponse})
}

func (c *CourseControllerImpl) EnrollCourse(ctx *gin.Context) {
	request := web.EnrollCourseInput{}

	courseId, err := strconv.Atoi(ctx.Param("courseId"))
	helper.PanicIfError(err)
	request.CourseId = courseId

	userId := ctx.MustGet("current_user").(web.UserResponse).Id
	request.UserId = userId

	enrollCourseResponse := c.CourseService.EnrollCourse(request)

	ctx.JSON(http.StatusOK, gin.H{"transaction": enrollCourseResponse})
}

func (c *CourseControllerImpl) UsersCompletedCourse(ctx *gin.Context) {
	user := ctx.MustGet("current_user").(web.UserResponse)
	courseId, err := strconv.Atoi(ctx.Param("courseId"))
	helper.PanicIfError(err)

	reponse, err := c.CourseService.IsCourseCompletedByUser(user.Id, courseId)
	helper.PanicIfError(err)

	ctx.JSON(200,
		helper.APIResponse(200, "Success to enrolled",
			gin.H{"enrolled_by": user.FirstName, "is_complete": reponse}),
	)
}

func NewCourseController(courseService service.CourseService) CourseController {
	return &CourseControllerImpl{CourseService: courseService}
}
