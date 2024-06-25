package service

import (
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/model/domain"
	"go-pzn-restful-api/model/web"
	"go-pzn-restful-api/repository"
	"os"
)

type LessonContentServiceImpl struct {
	repository.LessonContentRepository
	CourseService
	LessonService
	ChapterService
	repository.CourseRepository
}

func (s *LessonContentServiceImpl) FindById(lcId int) web.LessonContentResponse {
	lessonContent, err := s.LessonContentRepository.FindById(lcId)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	return helper.ToLessonContentResponse(lessonContent)
}

func (s *LessonContentServiceImpl) FindByLessonId(ltId int) []web.LessonContentResponse {
	lessonContents, err := s.LessonContentRepository.FindByLessonId(ltId)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	return helper.ToLessonContentsResponse(lessonContents)
}

func (s *LessonContentServiceImpl) Create(input web.ListLessonContentCreateInput) []web.LessonContentResponse {
	lessonContentsResponse := []web.LessonContentResponse{}

	lesson := s.LessonService.FindById(input.LessonId)

	chapter := s.ChapterService.FindById(lesson.ChapterId)

	course, err := s.CourseRepository.FindById(chapter.CourseId)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	if course.AuthorId != input.AuthorId {
		panic(helper.NewUnauthorizedError("You're not an author of this courses"))
	}

	for _, lessonContentInput := range input.LessonContents {
		lessonContent := domain.LessonContent{
			LessonId:     input.LessonId,
			Title:        lessonContentInput.Title,
			Content:      lessonContentInput.Content,
			Type:         lessonContentInput.Type,
			InOrder:      lessonContentInput.InOrder,
			Illustration: lessonContentInput.Illustration,
		}
		savedLessonContent := s.LessonContentRepository.Save(lessonContent)
		lessonContentResponse := helper.ToLessonContentResponse(savedLessonContent)
		lessonContentsResponse = append(lessonContentsResponse, lessonContentResponse)
	}

	lessonHash, err := helper.GenerateSHA256Hash(lessonContentsResponse)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	combinedHash, err := helper.GenerateSHA256Hash(course.HashCourse + lessonHash)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}
	course.HashCourse = combinedHash
	s.CourseRepository.Update(course)

	return lessonContentsResponse
}

func (s *LessonContentServiceImpl) UploadIllustration(lcId int, pathFile string) bool {
	findById, err := s.LessonContentRepository.FindById(lcId)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	if findById.Illustration != pathFile {
		if findById.Illustration == "" {
			return updateWhenUploadIllustration(findById, pathFile, s.LessonContentRepository)
		}
		os.Remove(findById.Illustration)
		return updateWhenUploadIllustration(findById, pathFile, s.LessonContentRepository)
	}

	return updateWhenUploadIllustration(findById, pathFile, s.LessonContentRepository)
}

func updateWhenUploadIllustration(lessonContent domain.LessonContent, pathFile string, lessonContentRepository repository.LessonContentRepository) bool {
	lessonContent.Illustration = pathFile
	lessonContentRepository.Update(lessonContent)
	return true
}

func NewLessonContentService(lessonContentRepository repository.LessonContentRepository, courseService CourseService, lessonService LessonService, chapterService ChapterService, courseRepository repository.CourseRepository) LessonContentService {
	return &LessonContentServiceImpl{
		LessonContentRepository: lessonContentRepository,
		CourseService:           courseService,
		LessonService:           lessonService,
		ChapterService:          chapterService,
		CourseRepository:        courseRepository,
	}
}
