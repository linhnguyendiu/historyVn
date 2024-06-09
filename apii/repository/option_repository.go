package repository

import "go-pzn-restful-api/model/domain"

type OptionRepository interface {
	Save(title domain.Option) domain.Option
	FindByQuestionId(questionId int) ([]domain.Option, error)
	FindById(opId int) (domain.Option, error)
	Update(title domain.Option) domain.Option
	GetCorrectOptionForQuestion(questionID int) (domain.Option, error)
}
