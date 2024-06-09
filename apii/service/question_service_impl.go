package service

import (
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/model/domain"
	"go-pzn-restful-api/model/web"
	"go-pzn-restful-api/repository"
)

type QuestionServiceImpl struct {
	repository.QuestionRepository
	CourseService
	repository.OptionRepository
}

func (s *QuestionServiceImpl) FindByCourseId(courseId int) []web.QuestionResponse {
	questions, err := s.QuestionRepository.FindByCourseId(courseId)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	questionsResponse := []web.QuestionResponse{}
	for _, question := range questions {
		options, err := s.OptionRepository.FindByQuestionId(question.Id)
		if err != nil {
			panic(helper.NewNotFoundError(err.Error()))
		}
		question.Option = options
		questionResponse := helper.ToQuestionResponse(question)
		questionsResponse = append(questionsResponse, questionResponse)
	}

	return questionsResponse
}

func (s *QuestionServiceImpl) Create(input web.QuestionCreateInput) web.QuestionResponse {
	ct := domain.Question{}
	ct.CourseId = input.CourseId
	ct.Content = input.Content

	course := s.CourseService.FindById(input.CourseId)
	if course.AuthorId != input.AuthorId {
		panic(helper.NewUnauthorizedError("You're not an author of this course"))
	}

	question := s.QuestionRepository.Save(ct)
	return helper.ToQuestionResponse(question)
}

func (s *QuestionServiceImpl) FindById(qsId int) web.QuestionResponse {
	question, err := s.QuestionRepository.FindQuestionById(qsId)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	options, err := s.OptionRepository.FindByQuestionId(qsId)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}
	question.Option = options

	return helper.ToQuestionResponse(question)
}

func NewQuestionService(questionRepository repository.QuestionRepository, courseService CourseService, optionRepository repository.OptionRepository) QuestionService {
	return &QuestionServiceImpl{
		QuestionRepository: questionRepository,
		CourseService:      courseService,
		OptionRepository:   optionRepository,
	}
}
