package web

import "time"

type CommentCreateInput struct {
	PostId           int `json:"post_id" binding:"required"`
	CommentFatherId  int `json:"comment_father_id"`
	UserId           int
	Content          string `json:"content" binding:"required"`
	CommentReply     bool   `json:"comment_reply"`
	ProfileImageAlt  string
	ProfileImageName string
}

type CommentResponse struct {
	Id               int               `json:"id"`
	UserId           int               `json:"user_id"`
	PostId           int               `json:"post_id"`
	CommentFatherId  int               `json:"comment_father_id"`
	Content          string            `json:"content"`
	Likes            int               `json:"likes"`
	Dislikes         int               `json:"dislikes"`
	Points           int               `json:"points"`
	CommentCount     int               `json:"comment_count"`
	ProfileImageAlt  string            `json:"profile_image_alt"`
	ProfileImageName string            `json:"profile_image_name"`
	CommentReply     bool              `json:"comment_reply"`
	CreatedAt        time.Time         `json:"created_at"`
	CommentChilds    []CommentResponse `json:"comment_childs"`
}
