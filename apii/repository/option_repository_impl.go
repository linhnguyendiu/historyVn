package repository

import (
	"errors"
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/model/domain"

	"gorm.io/gorm"
)

type OptionRepositoryImpl struct {
	db *gorm.DB
}

func (r *OptionRepositoryImpl) FindById(opId int) (domain.Option, error) {
	option := domain.Option{}
	err := r.db.Find(&option, "id=?", opId).Error
	if option.Id == 0 || err != nil {
		return option, errors.New("option not found")
	}

	return option, nil
}

func (r *OptionRepositoryImpl) Update(title domain.Option) domain.Option {
	err := r.db.Save(&title).Error
	helper.PanicIfError(err)

	return title
}

func (r *OptionRepositoryImpl) FindByQuestionId(questionId int) ([]domain.Option, error) {
	options := []domain.Option{}
	err := r.db.Find(&options, "question_id=?", questionId).Error
	if len(options) == 0 || err != nil {
		return nil, errors.New("option not found")
	}

	return options, nil
}

func (r *OptionRepositoryImpl) Save(title domain.Option) domain.Option {
	err := r.db.Create(&title).Error
	helper.PanicIfError(err)

	return title
}

func (r *OptionRepositoryImpl) GetCorrectOptionForQuestion(questionID int) (domain.Option, error) {
	option := domain.Option{}
	err := r.db.Where("question_id = ? AND is_correct = ?", questionID, true).First(&option).Error
	if option.Id == 0 || err != nil {
		return option, errors.New("option not found")
	}

	return option, nil
}

func NewOptionRepository(db *gorm.DB) OptionRepository {
	return &OptionRepositoryImpl{db: db}
}
