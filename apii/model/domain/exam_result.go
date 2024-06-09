package domain

import "time"

type ExamResult struct {
	CourseId           int       `gorm:"not null"`
	UserId             int       `gorm:"not null"`
	Score              int       `gorm:"not null"`
	SubmittedAt        time.Time `gorm:"type:TIMESTAMP;not null;default:CURRENT_TIMESTAMP"`
	TotalQuestions     int
	RewardAddress      string
	CertificateAddress string
}
