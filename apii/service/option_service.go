package service

import "go-pzn-restful-api/model/web"

type OptionService interface {
	Create(input web.OptionCreateInput) web.OptionResponse
	FindById(opId int) web.OptionResponse
	FindByQuestionId(ltId int) []web.OptionResponse
}
