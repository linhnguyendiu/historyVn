package web

import "time"

type RewardHistoryInput struct {
	UserId        int       `json:"user_id" binding:"required"`
	RewardType    int       `json:"reward_type" binding:"required"`
	RewardAt      time.Time `json:"reward_at" binding:"required"`
	RewardAddress string    `json:"reward_address" binding:"required"`
	CountReward   int       `json:"count_reward" binding:"required"`
}

type RewardHistoryResponse struct {
	UserId        int       `json:"user_id"`
	RewardType    int       `json:"reward_type"`
	RewardAt      time.Time `json:"reward_at"`
	RewardAddress string    `json:"reward_address"`
	CountReward   int       `json:"count_reward"`
}
