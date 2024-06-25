package repository

import (
	"errors"
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/model/domain"

	"gorm.io/gorm"
)

type ExamResultRepositoryImpl struct {
	db *gorm.DB
}

func (r *ExamResultRepositoryImpl) Save(examResult domain.ExamResult) domain.ExamResult {
	err := r.db.Create(&examResult).Error
	helper.PanicIfError(err)

	return examResult
}

func (r *ExamResultRepositoryImpl) Update(examResult domain.ExamResult) (domain.ExamResult, error) {
	var existingExamResult domain.ExamResult

	// Tìm bản ghi hiện tại trong cơ sở dữ liệu
	err := r.db.Where("course_id = ? AND user_id = ?", examResult.CourseId, examResult.UserId).First(&existingExamResult).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return domain.ExamResult{}, errors.New("exam result not found")
	} else if err != nil {
		return domain.ExamResult{}, err
	}

	// Cập nhật bản ghi
	err = r.db.Model(&existingExamResult).Where("course_id = ? AND user_id = ?", examResult.CourseId, examResult.UserId).Updates(examResult).Error
	if err != nil {
		return domain.ExamResult{}, err
	}

	return existingExamResult, nil
}

func (r *ExamResultRepositoryImpl) FindById(userId int, courseId int) (domain.ExamResult, error) {
	examResult := domain.ExamResult{}
	err := r.db.Where("user_id = ? AND course_id = ?", userId, courseId).First(&examResult).Error
	if examResult.UserId == 0 || err != nil {
		return examResult, errors.New("you not enroll this course")
	}

	return examResult, nil
}

func (r *ExamResultRepositoryImpl) HasUserEnrolledInCourse(userId, courseId int) (bool, error) {
	var count int64
	examResult := domain.ExamResult{}
	err := r.db.Model(&examResult).Where("user_id = ? AND course_id = ?", userId, courseId).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func NewExamResultRepository(db *gorm.DB) ExamResultRepository {
	return &ExamResultRepositoryImpl{db: db}
}
