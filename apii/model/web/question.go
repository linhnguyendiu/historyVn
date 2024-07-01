package web

type QuestionCreateInput struct {
	AuthorId int
	CourseId int
	Content  string           `json:"content"`
	Options  []OptionResponse `json:"options"`
}

type ListQuestionCreateInput struct {
	AuthorId  int
	CourseId  int
	Questions []QuestionResponse `json:"questions"`
}

type ListQuestionResponse struct {
	DurationQuiz int                `json:"duration_quiz"`
	QuizzesCount int                `json:"quizzes_count"`
	Questions    []QuestionResponse `json:"questions"`
}

type QuestionResponse struct {
	Id       int              `json:"id"`
	CourseId int              `json:"course_id"`
	Content  string           `json:"content"`
	Options  []OptionResponse `json:"options"`
}
