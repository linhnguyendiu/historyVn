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

	return imageCourse
}

func (r *ImageCourseRepositoryImpl) FindByCourseId(courseId int) ([]domain.ImageCourse, error) {
	images := []domain.ImageCourse{}
	err := r.db.Find(&images, "course_id=?", courseId).Error
	if len(images) == 0 || err != nil {
		return nil, errors.New("image not found")
	}

	return images, nil
}

func (r *ImageCourseRepositoryImpl) GetRandomImageByCourse(courseID int) (domain.ImageCourse, error) {
	var image domain.ImageCourse
	// randSource := rand.NewSource(time.Now().UnixNano())
	// rand.New(randSource)

	// Query the database for a random image of the given course and image type
	query := `SELECT * FROM image_courses WHERE course_id = ? ORDER BY RAND(), image_courses.id LIMIT 1`
	err := r.db.Raw(query, courseID).Scan(&image).Error
	// err := r.db.Where("course_id = ?", courseID).Order("RAND()").First(&image)
	if image.Id == 0 || err != nil {
		return image, errors.New("image not found")
	}

	return image, nil
}

func (r *ImageCourseRepositoryImpl) FindById(ctId int) (domain.ImageCourse, error) {
	image := domain.ImageCourse{}
	err := r.db.Find(&image, "id=?", ctId).Error
	if image.Id == 0 || err != nil {
		return image, errors.New("image not found")
	}

	return image, nil
}

func (r *ImageCourseRepositoryImpl) Update(course domain.ImageCourse) domain.ImageCourse {
	err := r.db.Save(&course).Error
	helper.PanicIfError(err)

	return course
}

func NewImageCourseRepository(db *gorm.DB) ImageCourseRepository {
	return &ImageCourseRepositoryImpl{db: db}
}
