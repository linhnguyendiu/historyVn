package repository

import (
	"errors"
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/model/domain"

	"gorm.io/gorm"
)

type RewardRepositoryImpl struct {
	db *gorm.DB
}

func (r *RewardRepositoryImpl) Save(reward domain.RewardHistory) domain.RewardHistory {
	err := r.db.Create(&reward).Error
	helper.PanicIfError(err)

	return reward
}

func (r *RewardRepositoryImpl) Update(reward domain.RewardHistory) domain.RewardHistory {
	err := r.db.Save(&reward).Error
	helper.PanicIfError(err)

	return reward
}

func (r *RewardRepositoryImpl) FindById(rewardId int) (domain.RewardHistory, error) {
	ar := domain.RewardHistory{}
	err := r.db.Find(&ar, "id=?", rewardId).Error
	if err != nil || ar.Id == 0 {
		return ar, errors.New("Transaction not found")
	}

	return ar, nil
}

func (r *RewardRepositoryImpl) FindByUserId(userId int) ([]domain.RewardHistory, error) {
	historys := []domain.RewardHistory{}
	err := r.db.Find(&historys, "user_id=?", userId).Error
	if len(historys) == 0 || err != nil {
		return nil, errors.New("historys not found")
	}

	return historys, nil
}

func NewRewardRepository(db *gorm.DB) RewardRepository {
	return &RewardRepositoryImpl{db: db}
}
