package controller

import (
	"fmt"
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/model/web"
	"go-pzn-restful-api/service"

	"github.com/gin-gonic/gin"
)

type UserControllerImpl struct {
	service.UserService
}

func (c *UserControllerImpl) Logout(ctx *gin.Context) {
	userId := ctx.MustGet("current_user").(web.UserResponse).Id
	userResponse := c.UserService.Logout(userId)
	if userResponse.Token == "" {
		ctx.JSON(200,
			helper.APIResponse(200, "You're logged out",
				gin.H{"user": userResponse.FirstName, "token": userResponse.Token}),
		)
	}
}

func (c *UserControllerImpl) UploadAvatar(ctx *gin.Context) {
	fileHeader, err := ctx.FormFile("avatar")
	helper.PanicIfError(err)
	userId := ctx.MustGet("current_user").(web.UserResponse).Id
	user := c.UserService.FindById(userId)
	filePath := fmt.Sprintf("assets/images/avatars/%s-%s", user.Email, f
	ileHeader.Filename)

	uploadAvatar := c.UserService.UploadAvatar(userId, filePath)

	err = ctx.SaveUploadedFile(fileHeader, filePath)
	helper.PanicIfError(err)

	ctx.JSON(
		200,
		helper.APIResponse(200, "Your avatar has been uploaded",
			gin.H{"user": uploadAvatar.FirstName, "avatar": uploadAvatar.Avatar, "is_uploaded": true}),
	)
}

func (c *UserControllerImpl) GetById(ctx *gin.Context) {
	// user_Id from token
	user := ctx.MustGet("current_user").(web.UserResponse)
	findById := c.UserService.FindById(user.Id)
	ctx.JSON(
		200,
		helper.APIResponse(200, "Current user: "+findById.FirstName, findById),
	)
}

func (c *UserControllerImpl) Login(ctx *gin.Context) {
	input := web.UserLoginInput{}
	err := ctx.ShouldBindJSON(&input)
	helper.PanicIfError(err)

	login := c.UserService.Login(input)

	ctx.JSON(
		200,
		helper.APIResponse(200, "You're logged in now", login),
	)
}

func (c *UserControllerImpl) Register(ctx *gin.Context) {
	input := web.UserRegisterInput{}
	err := ctx.ShouldBindJSON(&input)
	helper.PanicIfError(err)

	register := c.UserService.Register(input)
	ctx.JSON(200, helper.APIResponse(200, "Register is successfully", register))
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{UserService: userService}
}
