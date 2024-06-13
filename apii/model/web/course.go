package web

type CourseCreateInput struct {
	AuthorId     int
	Title        string `json:"title" binding:"required"`
	Slug         string `json:"slug" binding:"required"`
	Description  string `json:"description" binding:"required"`
	Price        int    `json:"price"`
	Reward       int    `json:"reward" binding:"required"`
	Category     string `json:"category" binding:"required"`
	DurationQuiz int
	QuizzesCount int
}

type CourseResponse struct {
	Id            int    `json:"id"`
	AuthorId      int    `json:"author_id"`
	Title         string `json:"title"`
	Slug          string `json:"slug"`
	Description   string `json:"description"`
	Price         int    `json:"price"`
	Reward        int    `json:"reward"`
	Banner        string `json:"banner"`
	Category      string `json:"category"`
	UsersEnrolled int    `json:"users_enrolled"`
	DurationQuiz  int
	QuizzesCount  int
}

type CourseBySlugResponse struct {
	Id            int            `json:"id"`
	AuthorId      int            `json:"author_id"`
	Title         string         `json:"title"`
	Slug          string         `json:"slug"`
	Description   string         `json:"description"`
	Price         int            `json:"price"`
	Reward        int            `json:"reward"`
	Banner        string         `json:"banner"`
	UsersEnrolled int            `json:"users_enrolled"`
	Author        AuthorResponse `json:"author"`
}
