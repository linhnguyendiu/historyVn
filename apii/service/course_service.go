package service

import (
	"context"
	"go-pzn-restful-api/model/domain"
	"go-pzn-restful-api/model/web"
)

type CourseService interface {
	Create(request web.CourseCreateInput) web.CourseResponse
	FindByKeyword(query string) ([]web.CourseResponse, error)
	FindById(courseId int) web.CourseResponse
	FindByType(typeCourse string) []web.CourseResponse
	FindByTypeAndCategory(typeCourse string, cateName string) []web.CourseResponse
	FindByAuthorId(authorId int) []web.CourseResponse
	FindByUserId(userId int) []web.CourseResponse
	FindByCategory(categoryName string) []web.CourseResponse
	FindAll() []web.CourseResponse
	UserEnrolled(userId int, courseId int) domain.UserCourse
	FindAllCourseIdByUserId(userId int) []string
	GetScore(ctx context.Context, request web.ExamRequest) web.ExamResultResponse
	FindTop3Coures() []web.CourseResponse
}
