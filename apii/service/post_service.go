package service

import (
	"context"
	"go-pzn-restful-api/model/web"
)

type PostService interface {
	Create(request web.PostCreateInput) web.PostBySearchResponse
	FindById(postId int) web.PostResponse
	FindByKeyword(keyword string) ([]web.PostBySearchResponse, error)
	FindByUserId(userId int) []web.PostBySearchResponse
	FindByTopic(topic string) []web.PostBySearchResponse
	FindAll() []web.PostBySearchResponse
	LikePost(ctx context.Context, userId int, postId int) (int, bool, error)
	DisLikePost(ctx context.Context, userId int, postId int) (int, bool, error)
	ProcessPosts() bool
	RewardPost(postId int, point int)
}
