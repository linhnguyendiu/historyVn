package repository

import (
	"errors"
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/model/domain"

	"gorm.io/gorm"
)

type LessonContentRepositoryImpl struct {
	db *gorm.DB
}

func (r *LessonContentRepositoryImpl) FindByLessonId(ltId int) ([]domain.LessonContent, error) {
	lessonContents := []domain.LessonContent{}
	err := r.db.Order("in_order asc").Find(&lessonContents, "lesson_id=?", ltId).Error
	if len(lessonContents) == 0 || err != nil {
		return nil, errors.New("lesson contents not found")
	}

	return lessonContents, nil
}

func (r *LessonContentRepositoryImpl) Update(content domain.LessonContent) domain.LessonContent {
	err := r.db.Save(&content).Error
	helper.PanicIfError(err)

	return content
}

func (r *LessonContentRepositoryImpl) FindById(lcId int) (domain.LessonContent, error) {
	lc := domain.LessonContent{}
	err := r.db.Find(&lc, "Id=?", lcId).Error
	if lc.Id == 0 || err != nil {
		return lc, errors.New("lesson content not found")
	}

	return lc, nil
}

func (r *LessonContentRepositoryImpl) Save(content domain.LessonContent) domain.LessonContent {
	err := r.db.Create(&content).Error
	helper.PanicIfError(err)

	return content
}

func NewLessonContentRepository(db *gorm.DB) LessonContentRepository {
	return &LessonContentRepositoryImpl{db: db}
}
