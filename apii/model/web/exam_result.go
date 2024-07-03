package web

import "time"

type ExamRequest struct {
	UserId      int
	CourseId    int
	SubmittedAt time.Time `json:"submitted_at"`
	Anwers      []int     `json:"anwers"`
}

type QuestionWithOptions struct {
	Content string `json:"content"`
}

type ExamResultResponse struct {
	CourseId       int       `json:"course_id"`
	UserId         int       `json:"user_id"`
	Score          int       `json:"score"`
	TotalQuestions int       `json:"total_questions"`
	Attempt        int       `json:"attempt"`
	HashAnswer     string    `json:"hash_answer"`
	SubmittedAt    time.Time `json:"submitted_at"`
	RewardAddress  string    `json:"reward_address"`
	CertificateId  int       `json:"certificate_id"`
}

type EnrollCourseInput struct {
	UserId     int
	CourseId   int
	EnrolledAt time.Time `json:"enrolled_at"`
}

type EnrollCourseResponse struct {
	UserId     int
	CourseId   int
	EnrolledAt time.Time `json:"enrolled_at"`
	Status     bool      `json:"status"`
}
