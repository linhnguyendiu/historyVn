package service

import "go-pzn-restful-api/model/web"

type RewardService interface {
	Create(input web.RewardHistoryInput) web.RewardHistoryResponse
	//Update(lcId int, input web.LessonContentUpdateInput) web.LessonContentResponse
	FindById(rId int) web.RewardHistoryResponse
	FindByUserId(rId int) []web.RewardHistoryResponse
}
