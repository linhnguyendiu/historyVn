package service

import "go-pzn-restful-api/model/web"

type LessonContentService interface {
	Create(input web.ListLessonContentCreateInput) []web.LessonContentResponse
	// Update(lcId int, input web.LessonContentUpdateInput) web.LessonContentResponse
	FindById(lcId int) web.LessonContentResponse
	FindByLessonId(ltId int) []web.LessonContentResponse
	UploadIllustration(courseId int, pathFile string) bool
}
