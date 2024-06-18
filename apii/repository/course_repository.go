package repository

import (
	"go-pzn-restful-api/model/domain"
)

type CourseRepository interface {
	Save(course domain.Course) domain.Course
	// SaveToCategoryCourse(categoryName string, courseId int) bool
	Update(course domain.Course) domain.Course
	FindById(courseId int) (domain.Course, error)
	FindByType(typeCourse string) ([]domain.Course, error)
	FindTop3Course(limit int) ([]domain.Course, error)
	FindByKeywords(keyword string, limit int) ([]domain.Course, error)
	FindByAuthorId(authorId int) ([]domain.Course, error)
	FindByUserId(userId int) ([]domain.Course, error)
	FindByCategory(categoryName string) ([]domain.Course, error)
	FindByTypeAndCategory(typeCourse string, cateName string) ([]domain.Course, error)
	FindAll() ([]domain.Course, error)
	UsersEnrolled(userCourse domain.UserCourse) domain.UserCourse
	CountUsersEnrolled(courseId int) int
	FindAllCourseIdByUserId(userId int) []string
	GetTotalQuestionsByCourseId(courseID int) (int64, error)
	CountTotalLessonsInCourse(courseID int) (int, error)
	SaveResult(examResult domain.ExamResult) domain.ExamResult
}
