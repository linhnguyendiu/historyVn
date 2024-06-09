package web

type LessonCreateInput struct {
	ChapterId    int
	CourseId     int
	Title        string `json:"title" binding:"required"`
	InOrder      int    `json:"in_order" binding:"required"`
	DurationTime int    `json:"duration_time" binding:"required"`
	AuthorId     int
}

type LessonResponse struct {
	Id           int    `json:"id"`
	ChapterId    int    `json:"chapter_id"`
	Title        string `json:"title"`
	InOrder      int    `json:"in_order"`
	DurationTime int    `json:"duration_time"`
}
