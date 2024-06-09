package domain

import "time"

type Option struct {
	Id         int `gorm:"primaryKey"`
	QuestionId int
	Content    string    `gorm:"not null;default:untitled"`
	IsCorrect  bool      `gorm:"not null`
	CreatedAt  time.Time `gorm:"type:TIMESTAMP;not null;default:CURRENT_TIMESTAMP"`
}
