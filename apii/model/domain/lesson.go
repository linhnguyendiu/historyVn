package domain

import "time"

type Lesson struct {
	Id            int `gorm:"primaryKey"`
	ChapterId     int
	Title         string `gorm:"not null;default:untitled"`
	DurationTime  int
	InOrder       int
	CreatedAt     time.Time `gorm:"type:TIMESTAMP;not null;default:CURRENT_TIMESTAMP"`
	LessonContent []LessonContent
}
