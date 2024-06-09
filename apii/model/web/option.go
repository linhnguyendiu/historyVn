package web

type OptionCreateInput struct {
	AuthorId   int
	CourseId   int
	QuestionId int
	Content    string `json:"content"`
	IsCorrect  bool   `json:"is_correct"`
}

type OptionResponse struct {
	Id         int    `json:"id"`
	QuestionId int    `json:"question_id"`
	Content    string `json:"title"`
	IsCorrect  bool   `json:"is_correct`
}
