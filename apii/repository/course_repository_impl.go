package repository

import (
	"errors"
	"fmt"
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/model/domain"

	"gorm.io/gorm"
)

type CourseRepositoryImpl struct {
	db *gorm.DB
}

func (r *CourseRepositoryImpl) FindAllCourseIdByUserId(userId int) []string {
	userCourses := []domain.UserCourse{}
	err := r.db.Select("course_id").Find(&userCourses, "user_id=?", userId).Error
	helper.PanicIfError(err)

	var allCourseId []string

	for _, userCourse := range userCourses {
		allCourseId = append(allCourseId, fmt.Sprintf("%d", userCourse.CourseId))
	}

	return allCourseId
}

func (r *CourseRepositoryImpl) FindByCategory(categoryName string) ([]domain.Course, error) {
	courses := []domain.Course{}
	err := r.db.Find(&courses, "category=?", categoryName).Error
	if len(courses) == 0 || err != nil {
		return nil, errors.New("course not found")
	}
	return courses, nil
}

func (r *CourseRepositoryImpl) FindByUserId(userId int) ([]domain.Course, error) {
	courses := []domain.Course{}
	err := r.db.
		Joins("JOIN user_courses ON user_courses.course_id=courses.id").
		Joins("JOIN users ON users.id=user_courses.user_id").
		Where("users.id=?", userId).
		Find(&courses).Error
	if len(courses) == 0 || err != nil {
		return nil, errors.New("courses not found")
	}
	return courses, nil
}

func (r *CourseRepositoryImpl) Update(course domain.Course) domain.Course {
	err := r.db.Save(&course).Error
	helper.PanicIfError(err)

	return course
}

func (r *CourseRepositoryImpl) CountUsersEnrolled(courseId int) int {
	var count int64
	userCourse := domain.UserCourse{}
	err := r.db.Find(&userCourse, "course_id=?", courseId).Count(&count).Error
	helper.PanicIfError(err)

	return int(count)
}

func (r *CourseRepositoryImpl) UsersEnrolled(userCourse domain.UserCourse) domain.UserCourse {
	err := r.db.Create(&userCourse).Error
	helper.PanicIfError(err)
	return userCourse
}

func (r *CourseRepositoryImpl) FindAll() ([]domain.Course, error) {
	courses := []domain.Course{}
	err := r.db.Find(&courses).Error
	if len(courses) == 0 || err != nil {
		return nil, errors.New("courses not found")
	}

	return courses, nil
}

func (r *CourseRepositoryImpl) FindByAuthorId(authorId int) ([]domain.Course, error) {
	courses := []domain.Course{}
	err := r.db.Find(&courses, "author_id=?", authorId).Error
	if err != nil || len(courses) == 0 {
		return nil, errors.New("courses not found")
	}
	return courses, nil
}

func (r *CourseRepositoryImpl) FindByType(typeCourse string) ([]domain.Course, error) {
	courses := []domain.Course{}
	err := r.db.Find(&courses, "type=?", typeCourse).Error
	if len(courses) == 0 || err != nil {
		return nil, errors.New("course not found")
	}
	return courses, nil
}

func (r *CourseRepositoryImpl) FindByTypeAndCategory(typeCourse string, cateName string) ([]domain.Course, error) {
	courses := []domain.Course{}
	err := r.db.Where("type = ? AND category = ?", typeCourse, cateName).Find(&courses).Error
	if len(courses) == 0 || err != nil {
		return nil, errors.New("course not found")
	}
	return courses, nil
}

func (r *CourseRepositoryImpl) FindById(courseId int) (domain.Course, error) {
	course := domain.Course{}
	err := r.db.Find(&course, "Id=?", courseId).Error
	if course.Id == 0 || err != nil {
		return course, errors.New("course not found")
	}

	return course, nil
}

func (r *CourseRepositoryImpl) FindByKeywords(keyword string, limit int) ([]domain.Course, error) {
	courses := []domain.Course{}
	if err := r.db.Where("title LIKE ?", "%"+keyword+"%").Limit(limit).Find(&courses).Error; err != nil {
		return nil, err
	}
	return courses, nil
}

func (r *CourseRepositoryImpl) FindTop3Course(limit int) ([]domain.Course, error) {
	courses := []domain.Course{}
	if err := r.db.Table("courses").
		Select("courses.*, COUNT(user_courses.course_id) as enroll_count").
		Joins("left join user_courses on user_courses.course_id = courses.id").
		Group("courses.id").
		Order("enroll_count DESC").
		Limit(limit).
		Find(&courses).Error; err != nil {
		return nil, err
	}
	return courses, nil
}

func (r *CourseRepositoryImpl) GetTotalQuestionsByCourseId(courseId int) (int64, error) {
	var count int64
	err := r.db.Model(&domain.Question{}).Where("course_id = ?", courseId).Count(&count).Error
	if err != nil {
		return 0, errors.New("course not found")
	}
	return count, nil
}

func (r *CourseRepositoryImpl) CountTotalLessonsInCourse(courseID int) (int, error) {
	var count int64
	if err := r.db.Model(&domain.Lesson{}).
		Joins("left join chapters on chapters.id = lessons.chapter_id").
		Where("chapters.course_id = ?", courseID).
		Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}

func (r *CourseRepositoryImpl) SaveResult(examResult domain.ExamResult) domain.ExamResult {
	err := r.db.Create(&examResult).Error
	helper.PanicIfError(err)

	return examResult
}

func (r *CourseRepositoryImpl) Save(course domain.Course) domain.Course {
	err := r.db.Create(&course).Error
	helper.PanicIfError(err)

	return course
}

func NewCourseRepository(db *gorm.DB) CourseRepository {
	return &CourseRepositoryImpl{db: db}
}
