package repository

import (
	"errors"
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/model/domain"

	"gorm.io/gorm"
)

type PostRepositoryImpl struct {
	db *gorm.DB
}

func (r *PostRepositoryImpl) FindByUserId(userId int) ([]domain.Post, error) {
	posts := []domain.Post{}
	err := r.db.Find(&posts, "user_id=?", userId).Error
	if len(posts) == 0 || err != nil {
		return nil, errors.New("post not found")
	}

	return posts, nil
}

func (r *PostRepositoryImpl) GetTotalCommentByPostId(postId int) (int64, error) {
	var count int64
	err := r.db.Model(&domain.Comment{}).Where("post_id = ?", postId).Count(&count).Error
	if err != nil {
		return 0, errors.New("post not found")
	}
	return count, nil
}

func (r *PostRepositoryImpl) FindAll() ([]domain.Post, error) {
	posts := []domain.Post{}
	err := r.db.Find(&posts).Error
	if len(posts) == 0 || err != nil {
		return nil, errors.New("posts not found")
	}

	return posts, nil
}

func (r *PostRepositoryImpl) Update(title domain.Post) domain.Post {
	err := r.db.Save(&title).Error
	helper.PanicIfError(err)

	return title
}

func (r *PostRepositoryImpl) FindByTopic(topic string) ([]domain.Post, error) {
	posts := []domain.Post{}
	err := r.db.Find(&posts, "topic=?", topic).Error
	if len(posts) == 0 || err != nil {
		return nil, errors.New("posts not found")
	}

	return posts, nil
}

func (r *PostRepositoryImpl) FindByKeywords(keyword string, limit int) ([]domain.Post, error) {
	posts := []domain.Post{}
	if err := r.db.Where("keyworks LIKE ?", "%"+keyword+"%").Limit(limit).Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *PostRepositoryImpl) FindById(postId int) (domain.Post, error) {
	post := domain.Post{}
	err := r.db.Find(&post, "Id=?", postId).Error
	if post.Id == 0 || err != nil {
		return post, errors.New("post not found")
	}

	return post, nil
}

func (r *PostRepositoryImpl) Save(title domain.Post) domain.Post {
	err := r.db.Create(&title).Error
	helper.PanicIfError(err)

	return title
}

func (r *PostRepositoryImpl) Delete(postId int) {
	post := domain.Post{}
	err := r.db.Where("id=?", postId).Delete(&post).Error

	helper.PanicIfError(err)
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &PostRepositoryImpl{db: db}
}
