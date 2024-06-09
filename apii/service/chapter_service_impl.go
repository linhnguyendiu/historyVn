package service

import (
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/model/domain"
	"go-pzn-restful-api/model/web"
	"go-pzn-restful-api/repository"
)

type ChapterServiceImpl struct {
	repository.ChapterRepository
	CourseService
}

func (s *ChapterServiceImpl) Update(ctId int, input web.ChapterCreateInput) web.ChapterResponse {
	findById, err := s.ChapterRepository.FindById(ctId)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}
	if input.Title != "" {
		findById.Title = input.Title
	}
	if input.InOrder != 0 {
		findById.InOrder = input.InOrder
	}

	return helper.ToChapterResponse(s.ChapterRepository.Update(findById))
}

func (s *ChapterServiceImpl) FindByCourseId(courseId int) []web.ChapterResponse {
	chapters, err := s.ChapterRepository.FindByCourseId(courseId)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	return helper.ToChaptersResponse(chapters)
}

func (s *ChapterServiceImpl) Create(input web.ChapterCreateInput) web.ChapterResponse {
	ct := domain.Chapter{}
	ct.CourseId = input.CourseId
	ct.Title = input.Title
	ct.InOrder = input.InOrder

	course := s.CourseService.FindById(input.CourseId)
	if course.AuthorId != input.AuthorId {
		panic(helper.NewUnauthorizedError("You're not an author of this course"))
	}

	chapter := s.ChapterRepository.Save(ct)
	return helper.ToChapterResponse(chapter)
}

func NewChapterService(chapterRepository repository.ChapterRepository, courseService CourseService) ChapterService {
	return &ChapterServiceImpl{
		ChapterRepository: chapterRepository,
		CourseService:     courseService,
	}
}
