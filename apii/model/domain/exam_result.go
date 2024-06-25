package domain

import "time"

type ExamResult struct {
	CourseId           int       `gorm:"not null"`
	UserId             int       `gorm:"not null"`
	EnrolledAt         time.Time `gorm:"type:TIMESTAMP;not null;default:CURRENT_TIMESTAMP"`
	Status             bool
	Score              int
	SubmittedAt        time.Time `gorm:"type:TIMESTAMP;default:CURRENT_TIMESTAMP"`
	TotalQuestions     int
	HashAnswer         string
	RewardAddress      string
	CertificateAddress string
}
