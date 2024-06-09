package repository

import (
	"errors"
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/model/domain"

	"gorm.io/gorm"
)

type ImageCourseRepositoryImpl struct {
	db *gorm.DB
}

func (r *ImageCourseRepositoryImpl) Save(imageCourse domain.ImageCourse) domain.ImageCourse {
	err := r.db.Create(&imageCourse).Error
	helper.PanicIfError(err)

	return category
}

func (r *ImageCourseRepositoryImpl) FindByCourseId(courseId int) ([]domain.ImageCourse, error) {
	chapters := []domain.Chapter{}
	err := r.db.Order("in_order asc").Find(&chapters, "course_id=?", courseId).Error
	if len(chapters) == 0 || err != nil {
		return nil, errors.New("chapter not found")
	}

	return chapters, nil
}

func (r *ImageCourseRepositoryImpl) FindById(ctId int) (domain.ImageCourse, error) {
	chapter := domain.Chapter{}
	err := r.db.Find(&chapter, "id=?", ctId).Error
	if chapter.Id == 0 || err != nil {
		return chapter, errors.New("chapter not found")
	}

	return chapter, nil
}

func NewImageCourseRepository(db *gorm.DB) ImageCourseRepository {
	return &ImageCourseRepositoryImpl{db: db}
}
