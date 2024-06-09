package repository

import (
	"errors"
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/model/domain"

	"gorm.io/gorm"
)

type LessonRepositoryImpl struct {
	db *gorm.DB
}

func (r *LessonRepositoryImpl) FindById(ltId int) (domain.Lesson, error) {
	lesson := domain.Lesson{}
	err := r.db.Find(&lesson, "id=?", ltId).Error
	if lesson.Id == 0 || err != nil {
		return lesson, errors.New("lesson not found")
	}

	return lesson, nil
}

func (r *LessonRepositoryImpl) Update(title domain.Lesson) domain.Lesson {
	err := r.db.Save(&title).Error
	helper.PanicIfError(err)

	return title
}

func (r *LessonRepositoryImpl) FindByChapterId(chapterId int) ([]domain.Lesson, error) {
	lessons := []domain.Lesson{}
	err := r.db.Order("in_order asc").Find(&lessons, "chapter_id=?", chapterId).Error
	if len(lessons) == 0 || err != nil {
		return nil, errors.New("lesson not found")
	}

	return lessons, nil
}

func (r *LessonRepositoryImpl) Save(title domain.Lesson) domain.Lesson {
	err := r.db.Create(&title).Error
	helper.PanicIfError(err)

	return title
}

func NewLessonRepository(db *gorm.DB) LessonRepository {
	return &LessonRepositoryImpl{db: db}
}
