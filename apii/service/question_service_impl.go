package service

import (
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/model/domain"
	"go-pzn-restful-api/model/web"
	"go-pzn-restful-api/repository"
	"log"
	"math/big"
)

type QuestionServiceImpl struct {
	repository.QuestionRepository
	CourseService
	repository.OptionRepository
	repository.CourseRepository
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

func (s *QuestionServiceImpl) CreateWithOptions(input web.ListQuestionCreateInput) []web.QuestionResponse {
	// Kiểm tra xem user có phải là tác giả của khóa học hay không
	questionsResponse := []web.QuestionResponse{}
	course, err := s.CourseRepository.FindById(input.CourseId)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}
	if course.AuthorId != input.AuthorId {
		panic(helper.NewUnauthorizedError("You're not an author of this course"))
	}

	for _, questionInput := range input.Questions {

		question := domain.Question{
			CourseId: input.CourseId,
			Content:  questionInput.Content,
		}
		savedQuestion := s.QuestionRepository.Save(question)
		// Tạo các tùy chọn
		options := []domain.Option{}
		for _, optionInput := range questionInput.Options {
			option := domain.Option{
				QuestionId: savedQuestion.Id,
				Content:    optionInput.Content,
				IsCorrect:  optionInput.IsCorrect,
			}
			savedOption := s.OptionRepository.Save(option)
			options = append(options, savedOption)
		}
		savedQuestion.Option = options

		// Chuyển đổi kết quả thành response
		questionResponse := helper.ToQuestionResponse(savedQuestion)
		questionsResponse = append(questionsResponse, questionResponse)
	}

	examHash, err := helper.GenerateSHA256Hash(questionsResponse)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	combinedHash, err := helper.GenerateSHA256Hash(course.HashCourse + examHash)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}
	course.HashCourse = combinedHash
	s.CourseRepository.Update(course)

	auth := helper.AuthGenerator(helper.Client)
	add, err := helper.Manage.AddCourse(auth, big.NewInt(int64(course.Id)), course.Title, big.NewInt(int64(course.Price)), big.NewInt(int64(course.Reward)), course.Type, course.HashCourse)
	if err != nil {
		helper.PanicIfError(err)
	}
	log.Printf("add successfull", add)

	totalQuestions, err := s.CourseRepository.GetTotalQuestionsByCourseId(course.Id)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}
	course.QuizzesCount = int(totalQuestions)

	s.CourseRepository.Update(course)

	return questionsResponse
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

func NewQuestionService(questionRepository repository.QuestionRepository, courseService CourseService, optionRepository repository.OptionRepository, courseRepository repository.CourseRepository) QuestionService {
	return &QuestionServiceImpl{
		QuestionRepository: questionRepository,
		CourseService:      courseService,
		OptionRepository:   optionRepository,
		CourseRepository:   courseRepository,
	}
}
