package repository

import "go-pzn-restful-api/model/domain"

type CommentRepository interface {
	Save(title domain.Comment) domain.Comment
	FindById(commentId int) (domain.Comment, error)
	FindByPostId(postId int) ([]domain.Comment, error)
	FindByUserId(postId int) ([]domain.Comment, error)
	FindByCommentFartherId(comFarId int) ([]domain.Comment, error)
	Update(title domain.Comment) domain.Comment
	Delete(commentId int)
	FindAll() ([]domain.Comment, error)
	CountCommentsByFatherId(commentId int) (int, error)
	CheckCommentExists(postId int, comFarId int) error
}
