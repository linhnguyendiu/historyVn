package repository

import "go-pzn-restful-api/model/domain"

type ImageCourseRepository interface {
	Save(imageCourse domain.ImageCourse) domain.ImageCourse
	FindByCourseId(courseId int) ([]domain.ImageCourse, error)
	FindById(imgId int) (domain.ImageCourse, error)
}
