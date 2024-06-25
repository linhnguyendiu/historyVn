package domain

import "time"

type Course struct {
	Id              int `gorm:"primaryKey"`
	AuthorId        int
	Title           string
	Type            string
	Description     string `gorm:"type:text"`
	Price           int    `gorm:"default:0;not null"`
	Reward          int    `gorm:"default:0;not null"`
	Category        string
	DurationQuiz    int
	DurationToLearn int
	QuizzesCount    int
	LessonsCount    int
	UsersEnrolled   int
	HashCourse      string
	CreatedAt       time.Time `gorm:"type:TIMESTAMP;not null;default:CURRENT_TIMESTAMP"`
	Users           []User    `gorm:"many2many:user_courses;"`
	Transaction     []Transaction
	Chapter         []Chapter
	Question        []Question
	ImageCourse     []ImageCourse
	Author          Author
}
