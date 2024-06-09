package web

type QuestionCreateInput struct {
	AuthorId int
	CourseId int
	Content  string `json:"content"`
}

type QuestionResponse struct {
	Id       int              `json:"id"`
	CourseId int              `json:"course_id"`
	Content  string           `json:"content"`
	Options  []OptionResponse `json:"options"`
}
