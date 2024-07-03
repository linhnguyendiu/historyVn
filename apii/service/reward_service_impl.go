package service

import (
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/model/domain"
	"go-pzn-restful-api/model/web"
	"go-pzn-restful-api/repository"
)

type RewardServiceImpl struct {
	repository.RewardRepository
}

func (s *RewardServiceImpl) FindById(rId int) web.RewardHistoryResponse {
	reward, err := s.RewardRepository.FindById(rId)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	return helper.ToRewardHistoryResponse(reward)
}

func (s *RewardServiceImpl) FindByUserId(userId int) []web.RewardHistoryResponse {
	rewards, err := s.RewardRepository.FindByUserId(userId)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	return helper.ToRewardHistorysResponse(rewards)
}

func (s *RewardServiceImpl) Create(input web.RewardHistoryInput) web.RewardHistoryResponse {
	option := domain.RewardHistory{}
	option.UserId = input.UserId
	option.RewardType = input.RewardType
	option.RewardAt = input.RewardAt
	option.RewardAddress = input.RewardAddress
	option.CountReward = input.CountReward

	content := s.RewardRepository.Save(option)
	//if err != nil {
	//	os.Remove(input.Content)
	//	helper.PanicIfError(err)
	//}

	return helper.ToRewardHistoryResponse(content)
}

func NewRewardService(rewardRepository repository.RewardRepository) RewardService {
	return &RewardServiceImpl{
		RewardRepository: rewardRepository,
	}
}
