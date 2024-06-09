package repository

import "go-pzn-restful-api/model/domain"

type ExamResultRepository interface {
	Save(examResult domain.ExamResult) domain.ExamResult
}
