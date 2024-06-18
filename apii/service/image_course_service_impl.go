package service

import (
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/model/domain"
	"go-pzn-restful-api/model/web"
	"go-pzn-restful-api/repository"
	"os"
)

type ImageCourseServiceImpl struct {
	repository.ImageCourseRepository
	repository.CourseRepository
}

func (s *ImageCourseServiceImpl) FindByCourseId(courseId int) []web.ImgCourseResponse {
	imgs, err := s.ImageCourseRepository.FindByCourseId(courseId)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	return helper.ToImgCoursesResponse(imgs)
}

func (s *ImageCourseServiceImpl) Create(input web.ImgCourseRequest) web.ImgCourseResponse {
	img := domain.ImageCourse{}
	img.CourseId = input.CourseId
	img.ImageType = input.ImageType
	img.Description = input.Description

	_, err := s.CourseRepository.FindById(img.CourseId)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	img = s.ImageCourseRepository.Save(img)
	return helper.ToImgCourseResponse(img)
}

func (s *ImageCourseServiceImpl) UploadImg(imgId int, pathFile string) bool {
	findById, err := s.ImageCourseRepository.FindById(imgId)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	if findById.ImageAlt != pathFile {
		if findById.ImageAlt == "" {
			return updateWhenUploadBanner(findById, pathFile, s.ImageCourseRepository)
		}
		os.Remove(findById.ImageAlt)
		return updateWhenUploadBanner(findById, pathFile, s.ImageCourseRepository)
	}

	return updateWhenUploadBanner(findById, pathFile, s.ImageCourseRepository)
}

func updateWhenUploadBanner(img domain.ImageCourse, pathFile string, imageCourseRepository repository.ImageCourseRepository) bool {
	img.ImageAlt = pathFile
	imageCourseRepository.Update(img)
	return true
}

func NewImageCourseService(imageCourseRepository repository.ImageCourseRepository, courseRepository repository.CourseRepository) ImageCourseService {
	return &ImageCourseServiceImpl{ImageCourseRepository: imageCourseRepository, CourseRepository: courseRepository}
}
