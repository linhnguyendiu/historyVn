package repository

import "go-pzn-restful-api/model/domain"

type PostRepository interface {
	Save(title domain.Post) domain.Post
	FindByUserId(userId int) ([]domain.Post, error)
	FindById(postId int) (domain.Post, error)
	Update(title domain.Post) domain.Post
	FindByTopic(topic string) ([]domain.Post, error)
	FindByKeywords(keyword string, limit int) ([]domain.Post, error)
	Delete(postId int)
	FindAll() ([]domain.Post, error)
	GetTotalCommentByPostId(postId int) (int64, error)
}
