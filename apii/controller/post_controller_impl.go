package controller

import (
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/model/web"
	"go-pzn-restful-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostControllerImpl struct {
	service.PostService
}

func (c *PostControllerImpl) GetByTopic(ctx *gin.Context) {
	postResponses := c.PostService.FindByTopic(ctx.Param("topicName"))

	ctx.JSON(200,
		helper.APIResponse(200, "List of posts", postResponses),
	)
}

// func (c *PostControllerImpl) UploadBanner(ctx *gin.Context) {
// 	courseStr := ctx.Param("courseId")
// 	courseId, _ := strconv.Atoi(courseStr)
// 	//courseId, _ := strconv.Atoi(ctx.Param("courseId"))

// 	fileHeader, _ := ctx.FormFile("banner")

// 	pathFile := fmt.Sprintf("assets/images/banners/%d-%s", courseId, fileHeader.Filename)
// 	uploadBanner := c.CourseService.UploadBanner(courseId, pathFile)

// 	ctx.SaveUploadedFile(fileHeader, pathFile)

// 	ctx.JSON(200,
// 		helper.APIResponse(200, "Banner is successfully uploaded",
// 			gin.H{"is_uploaded": uploadBanner}),
// 	)
// }

func (c *PostControllerImpl) GetAll(ctx *gin.Context) {
	postResponses := c.PostService.FindAll()
	ctx.JSON(200,
		helper.APIResponse(200, "List of posts", postResponses),
	)
}

func (c *PostControllerImpl) GetByUserId(ctx *gin.Context) {
	userId := ctx.MustGet("current_user").(web.UserResponse).Id

	postResponse := c.PostService.FindByUserId(userId)
	ctx.JSON(200,
		helper.APIResponse(200, "List of courses", postResponse),
	)
}

func (c *PostControllerImpl) GetByPostId(ctx *gin.Context) {
	postId, _ := strconv.Atoi(ctx.Param("postId"))
	postResponse := c.PostService.FindById(postId)
	ctx.JSON(200,
		helper.APIResponse(200, "post detail", postResponse),
	)
}

func (c *PostControllerImpl) GetByKeyword(ctx *gin.Context) {
	postResponses, err := c.PostService.FindByKeyword(ctx.Param("slug"))
	helper.PanicIfError(err)
	ctx.JSON(http.StatusOK,
		helper.APIResponse(200, "List of posts", postResponses),
	)
}

func (c *PostControllerImpl) Create(ctx *gin.Context) {
	request := web.PostCreateInput{}
	err := ctx.ShouldBindJSON(&request)
	helper.PanicIfError(err)

	userId := ctx.MustGet("current_user").(web.UserResponse).Id
	request.UserId = userId

	postResponse := c.PostService.Create(request)
	ctx.JSON(200,
		helper.APIResponse(200, "Post has been created", postResponse),
	)
}

func (c *PostControllerImpl) LikePost(ctx *gin.Context) {
	postId, _ := strconv.Atoi(ctx.Param("postId"))

	userId := ctx.MustGet("current_user").(web.UserResponse).Id

	likes, didUserLike, err := c.PostService.LikePost(ctx, userId, postId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":      "Post liked successfully",
		"likes":       likes,
		"didUserLike": didUserLike,
	})
}

func (c *PostControllerImpl) DisLikePost(ctx *gin.Context) {
	postId, _ := strconv.Atoi(ctx.Param("postId"))

	userId := ctx.MustGet("current_user").(web.UserResponse).Id

	disLikes, didUserDisLike, err := c.PostService.DisLikePost(ctx, userId, postId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":      "Post disliked successfully",
		"dislikes":    disLikes,
		"didUserLike": didUserDisLike,
	})
}

func (c *PostControllerImpl) ProcessPosts(ctx *gin.Context) {
	reponses := c.PostService.ProcessPosts()
	ctx.JSON(200,
		helper.APIResponse(200, "Calculate point all post successfull", reponses),
	)
}

func NewPostController(postService service.PostService) PostController {
	return &PostControllerImpl{PostService: postService}
}
