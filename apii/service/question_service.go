package service

import "go-pzn-restful-api/model/web"

type QuestionService interface {
	Create(title web.QuestionCreateInput) web.QuestionResponse
	FindByCourseId(courseId int) web.ListQuestionResponse
	FindById(qsId int) web.QuestionResponse
	CreateWithOptions(input web.ListQuestionCreateInput) []web.QuestionResponse
}
