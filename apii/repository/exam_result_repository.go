package repository

import "go-pzn-restful-api/model/domain"

type ExamResultRepository interface {
	Save(examResult domain.ExamResult) domain.ExamResult
	Update(examResult domain.ExamResult) (domain.ExamResult, error)
	FindById(userId int, courseId int) (domain.ExamResult, error)
	HasUserEnrolledInCourse(userId, courseId int) (bool, error)
}
