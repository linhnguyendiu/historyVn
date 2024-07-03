package domain

import "time"

type RewardHistory struct {
	Id            int `gorm:"primaryKey"`
	UserId        int `gorm:"not null"`
	RewardType    int
	RewardAt      time.Time `gorm:"type:TIMESTAMP;not null;default:CURRENT_TIMESTAMP"`
	RewardAddress string
	CountReward   int
}
