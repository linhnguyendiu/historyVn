package web

type ImgCourseRequest struct {
	CourseId    int
	ImageType   int    `json:"image_type" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type ImgCourseResponse struct {
	Id          int    `json:"id"`
	CourseId    int    `json:"course_id"`
	Description string `json:"description"`
	ImageType   int    `json:"image_type"`
	ImageAlt    string `json:"image_alt"`
}
