package domain

import "time"

type Comment struct {
	Id               int `gorm:"primaryKey"`
	PostId           int `gorm:"not null"`
	CommentFatherId  int
	UserId           int    `gorm:"not null"`
	Content          string `gorm:"type:text"`
	Likes            int
	Dislikes         int
	Points           int
	CommentReply     bool
	CommentCount     int
	RelatedPost      string
	User             User
	CommentChilds    []Comment `gorm:"-"`
	ProfileImageAlt  string
	ProfileImageName string
	CreatedAt        time.Time `gorm:"type:TIMESTAMP;not null;default:CURRENT_TIMESTAMP"`
}
