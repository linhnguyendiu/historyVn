package domain

import "time"

type Question struct {
	Id        int `gorm:"primaryKey"`
	CourseId  int
	Content   string    `gorm:"not null;default:untitled"`
	CreatedAt time.Time `gorm:"type:TIMESTAMP;not null;default:CURRENT_TIMESTAMP"`
	Option    []Option
}
