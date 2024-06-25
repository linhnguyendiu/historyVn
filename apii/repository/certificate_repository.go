package repository

import "go-pzn-restful-api/model/domain"

type CertificateRepository interface {
	Save(certificate domain.Certificate) domain.Certificate
	FindById(certId int) (domain.Certificate, error)
}
