package repository

import (
	"errors"
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/model/domain"

	"gorm.io/gorm"
)

type CommentRepositoryImpl struct {
	db *gorm.DB
}

func (r *CommentRepositoryImpl) Update(title domain.Comment) domain.Comment {
	err := r.db.Save(&title).Error
	helper.PanicIfError(err)

	return title
}

func (r *CommentRepositoryImpl) CountCommentsByFatherId(commentId int) (int, error) {
	var count int
	query := `
        WITH RECURSIVE CommentCTE AS (
            SELECT id
            FROM comments
            WHERE comment_father_id = ?
            UNION ALL
            SELECT c.id
            FROM comments c
            INNER JOIN CommentCTE ON c.comment_father_id = CommentCTE.id
        )
        SELECT COUNT(*)
        FROM CommentCTE;
    `
	err := r.db.Raw(query, commentId).Scan(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *CommentRepositoryImpl) FindAll() ([]domain.Comment, error) {
	comments := []domain.Comment{}
	err := r.db.Find(&comments).Error
	if len(comments) == 0 || err != nil {
		return nil, errors.New("comments not found")
	}

	return comments, nil
}

func (r *CommentRepositoryImpl) FindById(commentId int) (domain.Comment, error) {
	comment := domain.Comment{}
	err := r.db.Find(&comment, "Id=?", commentId).Error
	if comment.Id == 0 || err != nil {
		return comment, errors.New("comment not found")
	}
	return comment, nil
}

func (r *CommentRepositoryImpl) FindByPostId(postId int) ([]domain.Comment, error) {
	comments := []domain.Comment{}
	err := r.db.Find(&comments, "post_id=?", postId).Error
	if len(comments) == 0 || err != nil {
		return nil, errors.New("comments not found")
	}

	return comments, nil
}

func (r *CommentRepositoryImpl) FindByUserId(userId int) ([]domain.Comment, error) {
	comments := []domain.Comment{}
	err := r.db.Find(&comments, "user_id=?", userId).Error
	if len(comments) == 0 || err != nil {
		return nil, errors.New("comments not found")
	}

	return comments, nil
}

func (r *CommentRepositoryImpl) FindByCommentFartherId(comFarId int) ([]domain.Comment, error) {
	comments := []domain.Comment{}
	err := r.db.Find(&comments, "comment_father_id=?", comFarId).Error
	if len(comments) == 0 || err != nil {
		return nil, errors.New("comments child not found")
	}

	return comments, nil
}

func (r *CommentRepositoryImpl) CheckCommentExists(postId int, comFarId int) error {
	var count int64
	err := r.db.Model(&domain.Comment{}).Where("post_id = ? AND id = ?", postId, comFarId).Count(&count).Error
	if count == 0 || err != nil {
		return errors.New("comment father not exist in this post")
	}
	return nil
}

func (r *CommentRepositoryImpl) Save(title domain.Comment) domain.Comment {
	err := r.db.Create(&title).Error
	helper.PanicIfError(err)

	return title
}

func (r *CommentRepositoryImpl) Delete(commentId int) {
	comment := domain.Comment{}
	err := r.db.Where("id=?", commentId).Delete(&comment).Error

	helper.PanicIfError(err)
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &CommentRepositoryImpl{db: db}
}
