package web

type LessonContentCreateInput struct {
	AuthorId int
	CourseId int
	LessonId int
	Content  string `json:"content"`
	Type     int    `json:"type"`
	Title    string `json:"title"`
	InOrder  int    `json:"in_order"`
}

type LessonContentUpdateInput struct {
	AuthorId int
	CourseId int
	LessonId int
	Title    string `json:"title"`
	Content  string
	InOrder  int `json:"in_order"`
}

type LessonContentResponse struct {
	Id           int    `json:"id"`
	LessonId     int    `json:"lesson_id"`
	Title        string `json:"title"`
	Content      string `json:"content"`
	Type         int    `json:"type"`
	InOrder      int    `json:"in_order"`
	Illustration string `json:"illustration"`
}
