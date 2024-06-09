package domain

type TestResultDetail struct {
	CourseId   int  `gorm:"not null"`
	QuestionId int  `gorm:"not null"`
	OptionId   uint `gorm:"not null"`
}
