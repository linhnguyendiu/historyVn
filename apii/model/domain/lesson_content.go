package domain

import "time"

type LessonContent struct {
	Id           int `gorm:"primaryKey"`
	LessonId     int
	Description  string
	Content      string
	Illustration string
	InOrder      int       `gorm:"type:int"`
	CreatedAt    time.Time `gorm:"type:TIMESTAMP;not null;default:CURRENT_TIMESTAMP"`
}
