package app

import (
	"go-pzn-restful-api/auth"
	"go-pzn-restful-api/controller"
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/middleware"
	"go-pzn-restful-api/repository"
	"go-pzn-restful-api/service"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	jwtAuth = auth.NewJwtAuth()
	db      = DBConnection()

	// contract
	// contractRepository = repository.NewContractRepository(db, client, authh, linkToken, certNFT, eduManage)
	// contractService    = service.NewContractService(contractRepository)

	// user
	userRepository = repository.NewUserRepository(db)
	userService    = service.NewUserService(userRepository, jwtAuth)
	userController = controller.NewUserController(userService)

	// author
	authorRepository = repository.NewAuthorRepository(db)
	authorService    = service.NewAuthorService(authorRepository, jwtAuth)
	authorController = controller.NewAuthorController(authorService)

	// course
	courseRepository = repository.NewCourseRepository(db)
	courseService    = service.NewCourseService(courseRepository, optionRepository, examResultRepository)
	courseController = controller.NewCourseController(courseService)

	imageCourseRepository = repository.NewImageCourseRepository(db)
	imageCourseService    = service.NewImageCourseService(imageCourseRepository, courseRepository)
	imageCourseController = controller.NewImageCourseController(imageCourseService)

	// post
	postRepository = repository.NewPostRepository(db)
	postService    = service.NewPostService(postRepository, commentRepository, commentService, userService)
	postController = controller.NewPostController(postService)

	// comment
	commentRepository = repository.NewCommentRepository(db)
	commentService    = service.NewCommentService(commentRepository, postRepository, userService)
	commentController = controller.NewCommentController(commentService)

	// category
	categoryRepository = repository.NewCategoryRepository(db)
	categoryService    = service.NewCategoryService(categoryRepository)
	categoryController = controller.NewCategoryController(categoryService)

	// chapter
	chapterRepository = repository.NewChapterRepository(db)
	chapterService    = service.NewChapterService(chapterRepository, courseService, lessonRepository)
	chapterController = controller.NewChapterController(chapterService)

	// question
	questionRepository   = repository.NewQuestionRepository(db)
	examResultRepository = repository.NewExamResultRepository(db)
	questionService      = service.NewQuestionService(questionRepository, courseService, optionRepository)
	questionController   = controller.NewQuestionController(questionService)

	// option
	optionRepository = repository.NewOptionRepository(db)
	optionService    = service.NewOptionService(optionRepository, courseService)
	optionController = controller.NewOptionController(optionService)

	// lesson_title
	lessonRepository = repository.NewLessonRepository(db)
	lessonService    = service.NewLessonService(lessonRepository, courseService)
	lessonController = controller.NewLessonController(lessonService)

	// lesson_content
	lessonContentRepository = repository.NewLessonContentRepository(db)
	lessonContentService    = service.NewLessonContentService(lessonContentRepository, courseService)
	lessonContentController = controller.NewLessonContentController(lessonContentService)

	transactionController = controller.NewTransactionController(
		service.NewTransactionService(repository.NewTransactionRepository(db), courseService),
	)
)

func NewRouter() *gin.Engine {
	helper.EnvInit()
	helper.InitRedis()
	DBMigrate(db)

	// client := helper.DialClient()
	// helper.ConnectToLINKToken()
	// // helper.ConnectToCertNFT()
	// // helper.ConnectToEduManage()
	// auth := helper.AuthGenerator(client)
	// token := helper.GetTokenInstance()

	router := gin.Default()
	router.Use(cors.Default())
	router.Use(gin.CustomRecovery(middleware.ErrorHandler))

	v1 := router.Group("/api/v1")

	c := helper.NewCronHelper()

	// Lên lịch công việc
	_, err := c.AddFunc("*/30 * * * * ", func() {
		postService.ProcessPosts()
	})
	if err != nil {
		log.Println("Error adding cron job:", err)
		return nil
	}

	// Bắt đầu cron scheduler
	c.Start()

	// account := common.HexToAddress("0x5FbDB2315678afecb367f032d93F642f64180aa3")
	// value := big.NewInt(1000)

	// tranfer, err := token.Transfer(auth, account, value)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println("BalanceOf", tranfer)

	// balance, err := token.BalanceOf(&bind.CallOpts{}, account)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Println("BalanceOf", balance)

	// User endpoints
	v1.POST("/users", userController.Register)
	v1.POST("/users/login", userController.Login)
	v1.PATCH("/users/avatars", middleware.UserJwtAuthMiddleware(jwtAuth, userService), userController.UploadAvatar)
	v1.GET("/users", middleware.UserJwtAuthMiddleware(jwtAuth, userService), userController.GetById)
	v1.POST("/users/logout", middleware.UserJwtAuthMiddleware(jwtAuth, userService), userController.Logout)

	// Author endpoints
	v1.POST("/authors", authorController.Register)
	v1.GET("/authors", middleware.AuthorJwtAuthMiddleware(jwtAuth, authorService), authorController.GetById)
	v1.POST("/authors/login", authorController.Login)
	v1.PUT("/authors/avatars", middleware.AuthorJwtAuthMiddleware(jwtAuth, authorService), authorController.UploadAvatar)
	v1.POST("/authors/logout", middleware.AuthorJwtAuthMiddleware(jwtAuth, authorService), authorController.Logout)

	// Category endpoints
	v1.POST("/categories", middleware.AuthorJwtAuthMiddleware(jwtAuth, authorService), categoryController.Create)

	// Course endpoints
	v1.POST("/courses", middleware.AuthorJwtAuthMiddleware(jwtAuth, authorService), courseController.Create)
	// v1.PUT("/courses/:courseId/banners", middleware.AuthorJwtAuthMiddleware(jwtAuth, authorService), courseController.UploadBanner)
	v1.GET("/courses/authors/:authorId", courseController.GetByAuthorId)
	v1.GET("/courses/type/:type", courseController.GetByType)
	v1.GET("/courses", courseController.GetAll)
	v1.GET("/courses/special", courseController.GetTop3Course)
	v1.GET("/courses/categories/:categoryName", courseController.GetByCategory)
	v1.GET("/courses/:keywords", courseController.GetByKeyword)
	v1.GET("/courses/type/:type/categories/:cateName", courseController.GetByTypeAndCategory)
	v1.GET("/courses/enrolled", middleware.UserJwtAuthMiddleware(jwtAuth, userService), courseController.GetByUserId)
	// keknya dibawah ini ga perlu deh, soalnya udah ada transaksi endpoint
	v1.POST("/courses/:courseId/enrolled", middleware.UserJwtAuthMiddleware(jwtAuth, userService), courseController.UserEnrolled)

	v1.POST("/courses/:courseId/exam-result", middleware.AuthorJwtAuthMiddleware(jwtAuth, authorService), courseController.GetExamScore)

	v1.POST("/courses/:courseId/img", middleware.AuthorJwtAuthMiddleware(jwtAuth, authorService), imageCourseController.Create)
	v1.PUT("/courses/img/:imgId/imgAlt", middleware.AuthorJwtAuthMiddleware(jwtAuth, authorService), imageCourseController.UploadImg)
	v1.GET("/courses/img/:courseId", imageCourseController.GetByCourseId)

	// Post endpoints
	v1.POST("/posts", middleware.UserJwtAuthMiddleware(jwtAuth, userService), postController.Create)
	v1.GET("/posts/users", middleware.UserJwtAuthMiddleware(jwtAuth, userService), postController.GetByUserId)
	v1.GET("/posts", postController.GetAll)
	v1.GET("/posts/calculate-points", postController.ProcessPosts)
	v1.GET("/posts/postdetail/:postId", postController.GetByPostId)
	v1.GET("/posts/topics/:topicName", postController.GetByTopic)
	v1.GET("/posts/keywords/:slug", postController.GetByKeyword)
	v1.PATCH("/posts/likes/:postId", middleware.UserJwtAuthMiddleware(jwtAuth, userService), postController.LikePost)
	v1.PATCH("/posts/dislikes/:postId", middleware.UserJwtAuthMiddleware(jwtAuth, userService), postController.DisLikePost)

	// Comment endpoints
	v1.POST("/comments", middleware.UserJwtAuthMiddleware(jwtAuth, userService), commentController.Create)
	v1.GET("/comments", postController.GetAll)
	v1.GET("/comments/calculate-points", commentController.ProcessComments)
	v1.GET("/comments/comment-detail/:postId", commentController.GetCommentsByPostId)
	v1.GET("/comments/comment-list-detail/:commentId", commentController.GetByCommentFatherId)
	v1.PATCH("/comments/likes/:commentId", middleware.UserJwtAuthMiddleware(jwtAuth, userService), commentController.LikeComment)
	v1.PATCH("/comments/dislikes/:commentId", middleware.UserJwtAuthMiddleware(jwtAuth, userService), commentController.DisLikeComment)
	v1.GET("/comments/by-user", middleware.UserJwtAuthMiddleware(jwtAuth, userService), commentController.GetCommentsByUserId)

	// Chapter title endpoints
	v1.POST("/authors/courses/:courseId/chapter-titles", middleware.AuthorJwtAuthMiddleware(jwtAuth, authorService), chapterController.Create)
	v1.PATCH("/authors/courses/:courseId/chapter-titles/:ltId", middleware.AuthorJwtAuthMiddleware(jwtAuth, authorService), chapterController.Update)
	v1.GET("/courses/enrolled/:courseId/chapter-titles", chapterController.GetByCourseId)

	// Question title endpoints
	v1.POST("/authors/courses/:courseId/questions", middleware.AuthorJwtAuthMiddleware(jwtAuth, authorService), questionController.Create)
	v1.GET("/courses/enrolled/:courseId/questions", questionController.GetByCourseId)
	v1.GET("/courses/enrolled/:courseId/question/:questionId", questionController.GetByQuestionId)

	// Option title endpoints
	v1.POST("/authors/courses/:courseId/question/:questionId/option", middleware.AuthorJwtAuthMiddleware(jwtAuth, authorService), optionController.Create)
	v1.GET("/courses/enrolled/:courseId/question/:questionId/option", optionController.GetByQuestionId)

	// Lesson title endpoints
	v1.POST("/authors/courses/:courseId/chapter/:chapterId/lesson-titles", middleware.AuthorJwtAuthMiddleware(jwtAuth, authorService), lessonController.Create)
	v1.PATCH("/authors/courses/:courseId/lesson-titles/:ltId", middleware.AuthorJwtAuthMiddleware(jwtAuth, authorService), lessonController.Update)
	v1.GET("/lesson-complete/lesson-titles/:lessonId", middleware.UserJwtAuthMiddleware(jwtAuth, userService), lessonController.UsersCompletedLesson)
	v1.GET("/courses/enrolled/:courseId/chapter/:chapterId/lesson-titles", lessonController.GetByChapterId)

	// Lesson content endpoints
	v1.POST("/authors/courses/:courseId/lesson-titles/:ltId/lesson-contents", middleware.AuthorJwtAuthMiddleware(jwtAuth, authorService), lessonContentController.Create)
	v1.PUT("/authors/lesson-content/:lcId/illustrations", middleware.AuthorJwtAuthMiddleware(jwtAuth, authorService), lessonContentController.UploadIllustration)
	// v1.PATCH("/authors/courses/:courseId/lesson-contents/:lcId", middleware.AuthorJwtAuthMiddleware(jwtAuth, authorService), lessonContentController.Update)
	v1.GET("/c/:courseId/lesson-titles/:ltId/lesson-contents", lessonContentController.GetByLessonId)
	v1.GET("/c/:courseId/lesson-contents/:lcId",
		middleware.UserJwtAuthMiddleware(jwtAuth, userService),
		middleware.MidtransPaymentMiddleware(courseService),
		lessonContentController.GetById,
	)

	// transaction endpoints
	v1.POST("/courses/:courseId/transactions", middleware.UserJwtAuthMiddleware(jwtAuth, userService), transactionController.EarnPaidCourse)
	v1.POST("/transactions/notifications", transactionController.MidtransNotification)

	return router
}
