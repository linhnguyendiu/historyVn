package service

import "go-pzn-restful-api/model/web"

type LessonService interface {
	Create(title web.LessonCreateInput) web.LessonResponse
	FindByChapterId(chapterId int) []web.LessonResponse
	Update(ltId int, input web.LessonCreateInput) web.LessonResponse
}
