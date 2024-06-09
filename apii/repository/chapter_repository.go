package repository

import "go-pzn-restful-api/model/domain"

type ChapterRepository interface {
	Save(title domain.Chapter) domain.Chapter
	FindByCourseId(courseId int) ([]domain.Chapter, error)
	FindById(ctId int) (domain.Chapter, error)
	Update(title domain.Chapter) domain.Chapter
}
