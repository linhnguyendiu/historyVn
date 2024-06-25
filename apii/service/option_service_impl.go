package service

import (
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/model/domain"
	"go-pzn-restful-api/model/web"
	"go-pzn-restful-api/repository"
)

type OptionServiceImpl struct {
	repository.OptionRepository
	CourseService
}

func (s *OptionServiceImpl) FindById(lcId int) web.OptionResponse {
	option, err := s.OptionRepository.FindById(lcId)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	return helper.ToOptionResponse(option)
}

func (s *OptionServiceImpl) FindByQuestionId(qsId int) []web.OptionResponse {
	options, err := s.OptionRepository.FindByQuestionId(qsId)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	return helper.ToOptionsResponse(options)
}

	func (s *OptionServiceImpl) Create(input web.OptionCreateInput) web.OptionResponse {
		course := s.CourseService.FindById(input.CourseId)
		if course.AuthorId != input.AuthorId {
			panic(helper.NewUnauthorizedError("You're not an author of this courses"))
		}

		option := domain.Option{}
		option.QuestionId = input.QuestionId
		option.Content = input.Content
		option.IsCorrect = input.IsCorrect

		content := s.OptionRepository.Save(option)
		//if err != nil {
		//	os.Remove(input.Content)
		//	helper.PanicIfError(err)
		//}

		return helper.ToOptionResponse(content)
	}

func NewOptionService(optionRepository repository.OptionRepository, courseService CourseService) OptionService {
	return &OptionServiceImpl{
		OptionRepository: optionRepository,
		CourseService:    courseService,
	}
}
