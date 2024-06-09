package service

import (
	"context"
	"go-pzn-restful-api/model/web"
)

type CommentService interface {
	Create(request web.CommentCreateInput) web.CommentResponse
	FindById(commentId int) web.CommentResponse
	FindByComId(comFarId int) []web.CommentResponse
	FindByUserId(userId int) []web.CommentResponse
	FindByPostId(postId int) []web.CommentResponse
	LikeComment(ctx context.Context, userId int, commentId int) (int, bool, error)
	DisLikeComment(ctx context.Context, userId int, commentId int) (int, bool, error)
	ProcessComments() bool
}
