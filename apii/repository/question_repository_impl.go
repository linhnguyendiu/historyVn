package repository

import (
	"errors"
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/model/domain"

	"gorm.io/gorm"
)

type QuestionRepositoryImpl struct {
	db *gorm.DB
}

func (r *QuestionRepositoryImpl) FindQuestionById(qsId int) (domain.Question, error) {
	question := domain.Question{}
	err := r.db.Find(&question, "id=?", qsId).Error
	if question.Id == 0 || err != nil {
		return question, errors.New("question not found")
	}

	return question, nil
}

func (r *QuestionRepositoryImpl) Update(title domain.Question) domain.Question {
	err := r.db.Save(&title).Error
	helper.PanicIfError(err)

	return title
}

func (r *QuestionRepositoryImpl) FindByCourseId(courseId int) ([]domain.Question, error) {
	questions := []domain.Question{}
	err := r.db.Find(&questions, "course_id=?", courseId).Error
	if len(questions) == 0 || err != nil {
		return nil, errors.New("quiz not found")
	}

	return questions, nil
}

func (r *QuestionRepositoryImpl) Save(title domain.Question) domain.Question {
	err := r.db.Create(&title).Error
	helper.PanicIfError(err)

	return title
}

func NewQuestionRepository(db *gorm.DB) QuestionRepository {
	return &QuestionRepositoryImpl{db: db}
}
