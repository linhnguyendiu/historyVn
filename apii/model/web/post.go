package web

type PostCreateInput struct {
	UserId       int
	Title        string `json:"title" binding:"required"`
	Slug         string `json:"slug" binding:"required"`
	Description  string `json:"description" binding:"required"`
	Content      string `json:"content" binding:"required"`
	Topic        string `json:"topic" binding:"required"`
	Keyworks     string `json:"key_works" binding:"required"`
	CommentReply bool   `json:"comment_reply"`
}

type PostResponse struct {
	Id               int               `json:"id"`
	UserId           int               `json:"user_id"`
	Title            string            `json:"title"`
	Slug             string            `json:"slug"`
	Description      string            `json:"description"`
	Content          string            `json:"content"`
	Topic            string            `json:"topic"`
	Keyworks         string            `json:"key_works"`
	Banner           string            `json:"banner"`
	ProfileImageAlt  string            `json:"profile_image_alt"`
	ProfileImageName string            `json:"profile_image_name"`
	Likes            int               `json:"likes"`
	Dislikes         int               `json:"dislikes"`
	Points           int               `json:"points"`
	CommentCount     int64             `json:"comment_count"`
	CommentReply     bool              `json:"comment_reply"`
	Comments         []CommentResponse `json:"comments"`
}

type PostBySearchResponse struct {
	Id               int    `json:"id"`
	UserId           int    `json:"user_id"`
	Title            string `json:"title"`
	Slug             string `json:"slug"`
	Description      string `json:"description"`
	Content          string `json:"content"`
	Topic            string `json:"topic"`
	Keyworks         string `json:"key_works"`
	Banner           string `json:"banner"`
	Likes            int    `json:"likes"`
	Dislikes         int    `json:"dislikes"`
	Points           int    `json:"points"`
	CommentCount     int64  `json:"comment_count"`
	ProfileImageAlt  string `json:"profile_image_alt"`
	ProfileImageName string `json:"profile_image_name"`
}
