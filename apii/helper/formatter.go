package helper

import (
	"go-pzn-restful-api/model/domain"
	"go-pzn-restful-api/model/web"
)

func ToUserResponse(user domain.User) web.UserResponse {
	userResponse := web.UserResponse{}
	userResponse.Id = user.Id
	userResponse.LastName = user.LastName
	userResponse.FirstName = user.FirstName
	userResponse.Email = user.Email
	userResponse.Address = user.Address
	userResponse.Rank = user.Rank
	userResponse.Balance = user.Balance
	userResponse.Avatar = user.Avatar
	userResponse.Token = user.Token

	return userResponse
}

func ToAuthorResponse(author domain.Author) web.AuthorResponse {
	authorResponse := web.AuthorResponse{}
	authorResponse.Id = author.Id
	authorResponse.Name = author.Name
	authorResponse.Email = author.Email
	authorResponse.Profile = author.Profile
	authorResponse.Avatar = author.Avatar
	authorResponse.Token = author.Token

	return authorResponse
}

func ToCourseResponse(course domain.Course, countUserEnrolled int, countLessons int) web.CourseResponse {
	courseResponse := web.CourseResponse{}
	courseResponse.Id = course.Id
	courseResponse.AuthorId = course.AuthorId
	courseResponse.Title = course.Title
	courseResponse.Type = course.Type
	courseResponse.Description = course.Description
	courseResponse.Price = course.Price
	courseResponse.Reward = course.Reward
	courseResponse.Category = course.Category
	courseResponse.DurationQuiz = course.DurationQuiz
	courseResponse.UsersEnrolled = countUserEnrolled
	courseResponse.LessonsCount = countLessons
	return courseResponse
}

func ToPostResponse(post domain.Post) web.PostResponse {
	postResponse := web.PostResponse{}
	postResponse.Id = post.Id
	postResponse.UserId = post.UserId
	postResponse.Title = post.Title
	postResponse.Slug = post.Slug
	postResponse.Description = post.Description
	postResponse.Topic = post.Topic
	postResponse.Keyworks = post.Keyworks
	postResponse.Content = post.Content
	postResponse.ProfileImageAlt = post.ProfileImageAlt
	postResponse.ProfileImageName = post.ProfileImageName
	postResponse.Likes = post.Likes
	postResponse.Dislikes = post.Dislikes
	postResponse.Points = post.Points
	postResponse.CommentCount = post.CommentCount
	postResponse.Banner = post.Banner
	postResponse.CommentReply = post.CommentReply
	postResponse.Comments = ToCommentsResponse(post.Comments)

	return postResponse
}

func ToPostBySearchResponse(post domain.Post) web.PostBySearchResponse {
	postResponse := web.PostBySearchResponse{}
	postResponse.Id = post.Id
	postResponse.UserId = post.UserId
	postResponse.Title = post.Title
	postResponse.Slug = post.Slug
	postResponse.Description = post.Description
	postResponse.Topic = post.Topic
	postResponse.Keyworks = post.Keyworks
	postResponse.Content = post.Content
	postResponse.Likes = post.Likes
	postResponse.Dislikes = post.Dislikes
	postResponse.Points = post.Points
	postResponse.CommentCount = post.CommentCount
	postResponse.Banner = post.Banner
	postResponse.ProfileImageAlt = post.ProfileImageAlt
	postResponse.ProfileImageName = post.ProfileImageName

	return postResponse
}

func ToLessonContentResponse(content domain.LessonContent) web.LessonContentResponse {
	return web.LessonContentResponse{
		Id:           content.Id,
		LessonId:     content.LessonId,
		Title:        content.Title,
		Content:      content.Content,
		Illustration: content.Illustration,
		Type:         content.Type,
		InOrder:      content.InOrder,
	}
}

func ToLessonContentsResponse(contents []domain.LessonContent) []web.LessonContentResponse {
	lessonContents := []web.LessonContentResponse{}
	for _, content := range contents {
		lessonContents = append(lessonContents, ToLessonContentResponse(content))
	}

	return lessonContents
}

func ToCommentResponse(content domain.Comment) web.CommentResponse {
	commentResponse := web.CommentResponse{}
	commentResponse.Id = content.Id
	commentResponse.UserId = content.UserId
	commentResponse.PostId = content.PostId
	commentResponse.CommentFatherId = content.CommentFatherId
	commentResponse.Content = content.Content
	commentResponse.Likes = content.Likes
	commentResponse.Dislikes = content.Dislikes
	commentResponse.Points = content.Points
	commentResponse.CommentCount = content.CommentCount
	commentResponse.ProfileImageAlt = content.ProfileImageAlt
	commentResponse.ProfileImageName = content.ProfileImageName
	commentResponse.CommentReply = content.CommentReply
	commentResponse.CreatedAt = content.CreatedAt
	commentResponse.CommentChilds = ToCommentsResponse(content.CommentChilds)
	return commentResponse
}

func ToCommentsResponse(comments []domain.Comment) []web.CommentResponse {
	commentReponses := []web.CommentResponse{}
	for _, comment := range comments {
		commentReponses = append(commentReponses, ToCommentResponse(comment))
	}

	return commentReponses
}

func ToOptionResponse(option domain.Option) web.OptionResponse {
	return web.OptionResponse{
		Id:         option.Id,
		QuestionId: option.QuestionId,
		Content:    option.Content,
		IsCorrect:  option.IsCorrect,
	}
}

func ToOptionsResponse(options []domain.Option) []web.OptionResponse {
	optionsResponse := []web.OptionResponse{}
	for _, option := range options {
		optionResponse := ToOptionResponse(option)
		optionsResponse = append(optionsResponse, optionResponse)
	}
	return optionsResponse
}

// Convert domain Question to web QuestionResponse
func ToQuestionResponse(question domain.Question) web.QuestionResponse {
	return web.QuestionResponse{
		Id:       question.Id,
		CourseId: question.CourseId,
		Content:  question.Content,
		Options:  ToOptionsResponse(question.Option),
	}
}

// Convert domain Questions to web QuestionsResponse
func ToQuestionsResponse(questions []domain.Question) []web.QuestionResponse {
	questionResponses := []web.QuestionResponse{}
	for _, question := range questions {
		questionResponses = append(questionResponses, ToQuestionResponse(question))
	}
	return questionResponses
}

func ToMidtransTransactionResponse(transaction domain.Transaction, trxID string) web.MidtransTransactionResponse {
	return web.MidtransTransactionResponse{
		Id:         trxID,
		UserId:     transaction.UserId,
		CourseId:   transaction.CourseId,
		Amount:     transaction.Amount,
		Status:     transaction.Status,
		PaymentUrl: transaction.PaymentUrl,
	}
}

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToLessonResponse(title domain.Lesson) web.LessonResponse {
	return web.LessonResponse{
		Id:           title.Id,
		ChapterId:    title.ChapterId,
		Title:        title.Title,
		InOrder:      title.InOrder,
		DurationTime: title.DurationTime,
		Description:  title.Description,
		Type:         title.Type,
	}
}

func ToExamResultResponse(title domain.ExamResult) web.ExamResultResponse {
	return web.ExamResultResponse{
		CourseId:           title.CourseId,
		UserId:             title.UserId,
		Score:              title.Score,
		SubmittedAt:        title.SubmittedAt,
		TotalQuestions:     title.TotalQuestions,
		RewardAddress:      title.RewardAddress,
		CertificateAddress: title.CertificateAddress,
	}
}

func ToChapterResponse(title domain.Chapter) web.ChapterResponse {
	return web.ChapterResponse{
		Id:       title.Id,
		CourseId: title.CourseId,
		Title:    title.Title,
		InOrder:  title.InOrder,
		Lessons:  ToLessonsResponse(title.Lesson),
	}
}

func ToLessonsResponse(titles []domain.Lesson) []web.LessonResponse {
	lessonsResponse := []web.LessonResponse{}
	for _, lesson := range titles {
		lessonResponse := ToLessonResponse(lesson)
		lessonsResponse = append(lessonsResponse, lessonResponse)
	}

	return lessonsResponse
}

func ToChaptersResponse(titles []domain.Chapter) []web.ChapterResponse {
	chaptersResponse := []web.ChapterResponse{}
	for _, chapter := range titles {
		chapterResponse := ToChapterResponse(chapter)
		chaptersResponse = append(chaptersResponse, chapterResponse)
	}

	return chaptersResponse
}

func ToImgCourseResponse(title domain.ImageCourse) web.ImgCourseResponse {
	return web.ImgCourseResponse{
		Id:          title.Id,
		CourseId:    title.CourseId,
		Description: title.Description,
		ImageType:   title.ImageType,
		ImageAlt:    title.ImageAlt,
	}
}

func ToImgCoursesResponse(imgs []domain.ImageCourse) []web.ImgCourseResponse {
	imgsResponse := []web.ImgCourseResponse{}
	for _, img := range imgs {
		imgResponse := ToImgCourseResponse(img)
		imgsResponse = append(imgsResponse, imgResponse)
	}

	return imgsResponse
}
