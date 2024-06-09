package service

import (
	"context"
	"fmt"
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/model/domain"
	"go-pzn-restful-api/model/web"
	"go-pzn-restful-api/repository"
	"os"

	"github.com/go-redis/redis"
)

type CourseServiceImpl struct {
	repository.CourseRepository
	TransactionService
	repository.OptionRepository
	repository.ExamResultRepository
}

func (s *CourseServiceImpl) FindAllCourseIdByUserId(userId int) []string {
	return s.CourseRepository.FindAllCourseIdByUserId(userId)
}

func (s *CourseServiceImpl) FindByCategory(categoryName string) []web.CourseResponse {
	courses, err := s.CourseRepository.FindByCategory(categoryName)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	coursesResponse := []web.CourseResponse{}
	for _, course := range courses {
		countUsersEnrolled := s.CourseRepository.CountUsersEnrolled(course.Id)
		courseResponse := helper.ToCourseResponse(course, countUsersEnrolled)
		coursesResponse = append(coursesResponse, courseResponse)
	}

	return coursesResponse
}

func (s *CourseServiceImpl) FindByUserId(userId int) []web.CourseResponse {
	courses, err := s.CourseRepository.FindByUserId(userId)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	coursesResponse := []web.CourseResponse{}
	for _, course := range courses {
		countUsersEnrolled := s.CourseRepository.CountUsersEnrolled(course.Id)
		courseResponse := helper.ToCourseResponse(course, countUsersEnrolled)
		coursesResponse = append(coursesResponse, courseResponse)
	}

	return coursesResponse
}

func (s *CourseServiceImpl) UploadBanner(courseId int, pathFile string) bool {
	findById, err := s.CourseRepository.FindById(courseId)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	if findById.Banner != pathFile {
		if findById.Banner == "" {
			return updateWhenUploadBanner(findById, pathFile, s.CourseRepository)
		}
		os.Remove(findById.Banner)
		return updateWhenUploadBanner(findById, pathFile, s.CourseRepository)
	}

	return updateWhenUploadBanner(findById, pathFile, s.CourseRepository)
}

func (s *CourseServiceImpl) UserEnrolled(userId int, courseId int) domain.UserCourse {
	_, err := s.CourseRepository.FindById(courseId)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	userCourse := domain.UserCourse{
		CourseId: courseId,
		UserId:   userId,
	}

	usersEnrolled := s.CourseRepository.UsersEnrolled(userCourse)

	return usersEnrolled
}

func (s *CourseServiceImpl) FindAll() []web.CourseResponse {
	courses, err := s.CourseRepository.FindAll()
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	coursesResponse := []web.CourseResponse{}
	for _, course := range courses {
		countUsersEnrolled := s.CourseRepository.CountUsersEnrolled(course.Id)
		courseResponse := helper.ToCourseResponse(course, countUsersEnrolled)
		coursesResponse = append(coursesResponse, courseResponse)
	}

	return coursesResponse
}

func (s *CourseServiceImpl) FindByAuthorId(authorId int) []web.CourseResponse {
	courses, err := s.CourseRepository.FindByAuthorId(authorId)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	coursesResponse := []web.CourseResponse{}
	for _, course := range courses {
		countUsersEnrolled := s.CourseRepository.CountUsersEnrolled(course.Id)
		courseResponse := helper.ToCourseResponse(course, countUsersEnrolled)
		coursesResponse = append(coursesResponse, courseResponse)
	}

	return coursesResponse
}

func (s *CourseServiceImpl) FindBySlug(slug string) web.CourseBySlugResponse {
	findBySlug, err := s.CourseRepository.FindBySlug(slug)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	countUsersEnrolled := s.CourseRepository.CountUsersEnrolled(findBySlug.Id)
	return helper.ToCourseBySlugResponse(findBySlug, countUsersEnrolled)
}

func (s *CourseServiceImpl) FindById(courseId int) web.CourseResponse {
	findById, err := s.CourseRepository.FindById(courseId)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}
	//countUsersEnrolled := s.CourseRepository.CountUsersEnrolled(findById.Id)
	return helper.ToCourseResponse(findById, 1)
	//return helper.ToCourseResponse(findById, countUsersEnrolled)
}

func (s *CourseServiceImpl) Create(request web.CourseCreateInput) web.CourseResponse {
	course := domain.Course{
		AuthorId:    request.AuthorId,
		Title:       request.Title,
		Slug:        request.Slug,
		Description: request.Description,
		Price:       request.Price,
		Reward:      request.Reward,
	}

	if course.AuthorId == 0 {
		panic(helper.NewUnauthorizedError("You're not an author"))
	}

	save := s.CourseRepository.Save(course)
	// categoryCourse := s.CourseRepository.SaveToCategoryCourse(strings.ToLower(request.Category), save.Id)
	// if !categoryCourse {
	// 	panic(errors.New("Failed to create category for this course"))
	// }

	return helper.ToCourseResponse(save, 0)
}

func updateWhenUploadBanner(course domain.Course, pathFile string, courseRepository repository.CourseRepository) bool {
	course.Banner = pathFile
	courseRepository.Update(course)
	return true
}

func (s *CourseServiceImpl) GetScore(ctx context.Context, request web.ExamRequest) web.ExamResultResponse {
	examResult := domain.ExamResult{
		CourseId:       request.CourseId,
		UserId:         request.UserId,
		Score:          0,
		TotalQuestions: 0,
	}

	totalQuestions, err := s.CourseRepository.GetTotalQuestionsByCourseId(examResult.CourseId)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}
	examResult.TotalQuestions = int(totalQuestions)

	for _, answerID := range request.Anwers {
		option, err := s.OptionRepository.FindById(answerID)
		if err != nil {
			panic(helper.NewNotFoundError(err.Error()))
		}
		if option.IsCorrect {
			examResult.Score++
		}
	}
	examResult.Score = int((float64(examResult.Score) / float64(examResult.TotalQuestions)) * 10)
	//save := s.ExamResultRepository.Save(examResult)

	// Kiểm tra lượt làm của người dùng
	userAttemptsKey := fmt.Sprintf("user:%d:course:%d:attempts", request.UserId, request.CourseId)
	attemptCount, err := helper.RedisCli.Get(ctx, userAttemptsKey).Int()
	if err == redis.Nil {
		attemptCount = 0
	}
	// } else if err != nil {
	// 	panic(err)
	// }

	attemptCount++

	if attemptCount == 1 {
		// Lần làm đầu tiên, lưu vào cơ sở dữ liệu
		save := s.ExamResultRepository.Save(examResult)
		if err := helper.RedisCli.Set(ctx, userAttemptsKey, 1, 0).Err(); err != nil {
			panic(err)
		}
		return helper.ToExamResultResponse(save)
	} else {
		// Lần làm thứ 2 trở đi, lưu vào Redis
		scoreKey := fmt.Sprintf("user:%d:course:%d:attempt:%d:score", request.UserId, request.CourseId, attemptCount)
		if err := helper.RedisCli.Set(ctx, scoreKey, examResult.Score, 0).Err(); err != nil {
			panic(err)
		}
		if err := helper.RedisCli.Set(ctx, userAttemptsKey, attemptCount, 0).Err(); err != nil {
			panic(err)
		}
		return web.ExamResultResponse{
			UserId:         request.UserId,
			CourseId:       request.CourseId,
			Attempt:        attemptCount,
			Score:          examResult.Score,
			TotalQuestions: examResult.TotalQuestions,
		}
	}
}

func NewCourseService(courseRepository repository.CourseRepository, optionRepository repository.OptionRepository, examResultRepository repository.ExamResultRepository) CourseService {
	return &CourseServiceImpl{CourseRepository: courseRepository, OptionRepository: optionRepository, ExamResultRepository: examResultRepository}
}
