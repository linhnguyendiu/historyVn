package middleware

import (
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/model/web"
	"go-pzn-restful-api/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func MidtransPaymentMiddleware(courseService service.CourseService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID := ctx.MustGet("current_user").(web.UserResponse).Id
		strCourseID := ctx.Param("courseId")
		courseID, _ := strconv.Atoi(strCourseID)

		isUserHas := false

		checkUserEnrollment, err := courseService.CheckUserEnrollment(userID, courseID)
		if err != nil {
			panic(helper.NewNotFoundError(err.Error()))
		}
		if !checkUserEnrollment {
			panic(helper.NewNotFoundError("User has not enrolled the course"))
		}

		ctx.Set("isUserHas", isUserHas)
	}
}
