package repository

import "go-pzn-restful-api/model/domain"

type CertificateRepository interface {
	Save(certificate domain.Certificate) domain.Certificate
	FindById(certId int, userId int) (domain.Certificate, error)
}
