package controller

import (
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/model/web"
	"go-pzn-restful-api/service"

	"github.com/gin-gonic/gin"
)

type RewardControllerImpl struct {
	service.RewardService
}

func (c *RewardControllerImpl) GetByUserId(ctx *gin.Context) {
	userId := ctx.MustGet("current_user").(web.UserResponse).Id
	rewardsResponse := c.RewardService.FindByUserId(userId)

	ctx.JSON(200,
		helper.APIResponse(200, "List of rewards", rewardsResponse))
}

func (c *RewardControllerImpl) Create(ctx *gin.Context) {
	input := web.RewardHistoryInput{}
	err := ctx.ShouldBindJSON(&input)
	helper.PanicIfError(err)

	userId := ctx.MustGet("current_user").(web.UserResponse).Id
	input.UserId = userId

	rewardResponse := c.RewardService.Create(input)
	ctx.JSON(200,
		helper.APIResponse(200, "reward is successfully created", rewardResponse))
}

func NewRewardController(rewardService service.RewardService) RewardController {
	return &RewardControllerImpl{RewardService: rewardService}
}
