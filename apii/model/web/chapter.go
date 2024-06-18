package web

type ChapterCreateInput struct {
	CourseId int
	Title    string `json:"title" binding:"required"`
	InOrder  int    `json:"in_order" binding:"required"`
	AuthorId int
}

type ChapterResponse struct {
	Id       int              `json:"id"`
	CourseId int              `json:"course_id"`
	Title    string           `json:"title"`
	InOrder  int              `json:"in_order"`
	Lessons  []LessonResponse `json:"lessons"`
}
