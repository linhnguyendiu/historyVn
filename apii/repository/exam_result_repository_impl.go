package repository

import (
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/model/domain"

	"gorm.io/gorm"
)

type ExamResultRepositoryImpl struct {
	db *gorm.DB
}

func (r *ExamResultRepositoryImpl) Save(examResult domain.ExamResult) domain.ExamResult {
	err := r.db.Create(&examResult).Error
	helper.PanicIfError(err)

	return examResult
}

func NewExamResultRepository(db *gorm.DB) ExamResultRepository {
	return &ExamResultRepositoryImpl{db: db}
}
