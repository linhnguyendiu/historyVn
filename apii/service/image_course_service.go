package service

import "go-pzn-restful-api/model/web"

type ImageCourseService interface {
	Create(title web.ImgCourseRequest) web.ImgCourseResponse
	FindByCourseId(courseId int) []web.ImgCourseResponse
	UploadImg(imgId int, pathFile string) bool
}
