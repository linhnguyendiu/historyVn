package web

import "time"

type CertificateCreateInput struct {
	UserName   string `json:"user_name"`
	UserId     int
	CourseId   int
	CourseName string    `json:"course_name"`
	Date       time.Time `json:"date"`
	CertType   string    `json:"cert_type"`
	ImageUri   string    `json:"image_uri"`
}

type CertificateResponse struct {
	Id         int `json:"id"`
	UserId     int
	CourseId   int
	UserName   string    `json:"user_name"`
	CourseName string    `json:"course_name"`
	Date       time.Time `json:"date"`
	CertType   string    `json:"cert_type"`
	ImageUri   string    `json:"image_uri"`
	CertUri    string    `json:"cert_uri"`
}
