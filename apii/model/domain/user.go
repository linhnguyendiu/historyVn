package domain

import "time"

type User struct {
	Id          int `gorm:"primaryKey"`
	FirstName   string
	LastName    string
	Email       string
	Password    string
	Address     string
	Avatar      string
	Balance     int
	Rank        int
	Token       string
	CreatedAt   time.Time `gorm:"type:TIMESTAMP;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `gorm:"type:TIMESTAMP;not null;default:CURRENT_TIMESTAMP"`
	Courses     []Course  `gorm:"many2many:user_courses;"`
	Transaction []Transaction
}
