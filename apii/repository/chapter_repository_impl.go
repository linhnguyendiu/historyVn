package repository

import (
	"errors"
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/model/domain"

	"gorm.io/gorm"
)

type ChapterRepositoryImpl struct {
	db *gorm.DB
}

func (r *ChapterRepositoryImpl) FindById(ctId int) (domain.Chapter, error) {
	chapter := domain.Chapter{}
	err := r.db.Find(&chapter, "id=?", ctId).Error
	if chapter.Id == 0 || err != nil {
		return chapter, errors.New("chapter not found")
	}

	return chapter, nil
}

func (r *ChapterRepositoryImpl) Update(title domain.Chapter) domain.Chapter {
	err := r.db.Save(&title).Error
	helper.PanicIfError(err)

	return title
}

func (r *ChapterRepositoryImpl) FindByCourseId(courseId int) ([]domain.Chapter, error) {
	chapters := []domain.Chapter{}
	err := r.db.Order("in_order asc").Find(&chapters, "course_id=?", courseId).Error
	if len(chapters) == 0 || err != nil {
		return nil, errors.New("chapter not found")
	}

	return chapters, nil
}

func (r *ChapterRepositoryImpl) Save(title domain.Chapter) domain.Chapter {
	err := r.db.Create(&title).Error
	helper.PanicIfError(err)

	return title
}

func NewChapterRepository(db *gorm.DB) ChapterRepository {
	return &ChapterRepositoryImpl{db: db}
}
