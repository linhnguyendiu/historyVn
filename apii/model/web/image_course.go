package web

type ImgCourseRequest struct {
	CourseId  int
	ImageType string `json:"image_type"`
}

type ImgCourseResponse struct {
	Id        int    `json:"id"`
	CourseId  int    `json:"course_id"`
	ImageType string `json:"image_type"`
	ImageAlt  string `json:"image_alt"`
}
