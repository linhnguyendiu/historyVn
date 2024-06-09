package domain

import "time"

type Post struct {
	Id               int `gorm:"primaryKey"`
	UserId           int
	Title            string `gorm:"type:text"`
	Slug             string `gorm:"type:text"`
	Description      string `gorm:"type:text"`
	Content          string `gorm:"default:0;not null"`
	Topic            string `gorm:"default:0;not null"`
	Keyworks         string `gorm:"type:text"`
	Banner           string
	Likes            int
	Dislikes         int
	Points           int
	CommentCount     int64
	CommentReply     bool
	ProfileImageAlt  string
	ProfileImageName string
	Comments         []Comment `gorm:"-"`
	User             User
	CreatedAt        time.Time `gorm:"type:TIMESTAMP;not null;default:CURRENT_TIMESTAMP"`
}
