package repository

import "go-pzn-restful-api/model/domain"

type QuestionRepository interface {
	Save(title domain.Question) domain.Question
	FindByCourseId(courseId int) ([]domain.Question, error)
	FindQuestionById(qsId int) (domain.Question, error)
	Update(title domain.Question) domain.Question
}
