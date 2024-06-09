package repository

import "go-pzn-restful-api/model/domain"

type LessonRepository interface {
	Save(title domain.Lesson) domain.Lesson
	FindByChapterId(chapterId int) ([]domain.Lesson, error)
	FindById(ltId int) (domain.Lesson, error)
	Update(title domain.Lesson) domain.Lesson
}
