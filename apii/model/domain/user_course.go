package domain

import "time"

type UserCourse struct {
	CourseId  int
	UserId    int
	CreatedAt time.Time `gorm:"type:TIMESTAMP;not null;default:CURRENT_TIMESTAMP"`
}
