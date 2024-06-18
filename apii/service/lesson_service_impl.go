package service

import (
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/model/domain"
	"go-pzn-restful-api/model/web"
	"go-pzn-restful-api/repository"
)

type LessonServiceImpl struct {
	repository.LessonRepository
	CourseService
}

func (s *LessonServiceImpl) Update(ltId int, input web.LessonCreateInput) web.LessonResponse {
	findById, err := s.LessonRepository.FindById(ltId)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}
	if input.Title != "" {
		findById.Title = input.Title
	}
	if input.InOrder != 0 {
		findById.InOrder = input.InOrder
	}

	return helper.ToLessonResponse(s.LessonRepository.Update(findById))
}

func (s *LessonServiceImpl) FindByChapterId(chapterId int) []web.LessonResponse {
	lessons, err := s.LessonRepository.FindByChapterId(chapterId)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	return helper.ToLessonsResponse(lessons)
}

func (s *LessonServiceImpl) UsersCompletedLesson(userId int, lessonId int) domain.UserLesson {
	_, err := s.LessonRepository.FindById(lessonId)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	userLesson := domain.UserLesson{
		LessonId: lessonId,
		UserId:   userId,
	}

	usersCompletedLesson := s.LessonRepository.UsersCompletedLesson(userLesson)

	return usersCompletedLesson
}

func (s *LessonServiceImpl) Create(input web.LessonCreateInput) web.LessonResponse {
	lt := domain.Lesson{}
	lt.ChapterId = input.ChapterId
	lt.Title = input.Title
	lt.InOrder = input.InOrder
	lt.DurationTime = input.DurationTime
	lt.Description = input.Description
	lt.Type = input.Type

	course := s.CourseService.FindById(input.CourseId)
	if course.AuthorId != input.AuthorId {
		panic(helper.NewUnauthorizedError("You're not an author of this course"))
	}

	lesson := s.LessonRepository.Save(lt)
	return helper.ToLessonResponse(lesson)
}

func NewLessonService(lessonRepository repository.LessonRepository, courseService CourseService) LessonService {
	return &LessonServiceImpl{
		LessonRepository: lessonRepository,
		CourseService:    courseService,
	}
}
