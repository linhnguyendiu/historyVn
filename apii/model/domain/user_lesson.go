package domain

import "time"

type UserLesson struct {
	LessonId    int
	UserId      int
	CompletedAt time.Time `gorm:"type:TIMESTAMP;not null;default:CURRENT_TIMESTAMP"`
}
