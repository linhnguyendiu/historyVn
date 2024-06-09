package web

type LessonContentCreateInput struct {
	AuthorId    int
	CourseId    int
	LessonId    int
	Content     string `json:"content"`
	Description string `json:"description"`
	InOrder     int    `json:"in_order"`
}

type LessonContentUpdateInput struct {
	AuthorId    int
	CourseId    int
	LessonId    int
	Description string `json:"description"`
	Content     string
	InOrder     int `json:"in_order"`
}

type LessonContentResponse struct {
	Id           int    `json:"id"`
	LessonId     int    `json:"lesson_id"`
	Content      string `json:"content"`
	Description  string `json:"description"`
	InOrder      int    `json:"in_order"`
	Illustration string `json:"illustration"`
}
