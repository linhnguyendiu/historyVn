package domain

import "time"

type Certificate struct {
	Id             int       `gorm:"primaryKey"`
	CourseId       int       `gorm:"not null"`
	UserId         string    `gorm:"not null"`
	Score          int       `gorm:"not null"`
	SubmittedAt    time.Time `gorm:"type:TIMESTAMP;not null;default:CURRENT_TIMESTAMP"`
	RewardAdd      string
	CertificateAdd string
}
