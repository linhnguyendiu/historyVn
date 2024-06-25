package web

type CourseCreateInput struct {
	AuthorId     int
	Title        string `json:"title" binding:"required"`
	Type         string `json:"type" binding:"required"`
	Description  string `json:"description" binding:"required"`
	Price        int    `json:"price"`
	Reward       int    `json:"reward" binding:"required"`
	Category     string `json:"category" binding:"required"`
	DurationQuiz int    `json:"duration_quiz" binding:"required"`
	QuizzesCount int
	LessonsCount int
	ImageCourses []ImgCourseRequest   `json:"image_courses"`
	Chapters     []ChapterCreateInput `json:"chapters"`
}

type CourseResponse struct {
	Id              int    `json:"id"`
	AuthorId        int    `json:"author_id"`
	Title           string `json:"title"`
	Type            string `json:"type"`
	Category        string `json:"category"`
	Description     string `json:"description"`
	Price           int    `json:"price"`
	Reward          int    `json:"reward"`
	UsersEnrolled   int    `json:"users_enrolled"`
	DurationQuiz    int    `json:"duration_quiz"`
	DurationToLearn int
	LessonsCount    int
	QuizzesCount    int
	HashCourse      string
}
