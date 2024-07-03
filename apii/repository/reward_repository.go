package repository

import "go-pzn-restful-api/model/domain"

type RewardRepository interface {
	Save(reward domain.RewardHistory) domain.RewardHistory
	Update(reward domain.RewardHistory) domain.RewardHistory
	FindById(rewardId int) (domain.RewardHistory, error)
	FindByUserId(userId int) ([]domain.RewardHistory, error)
}
