package service

import "go-pzn-restful-api/model/web"

type ChapterService interface {
	Create(title web.ChapterCreateInput) web.ChapterResponse
	FindByCourseId(courseId int) []web.ChapterResponse
	Update(ctId int, input web.ChapterCreateInput) web.ChapterResponse
	FindById(lessonId int) web.ChapterResponse
}
