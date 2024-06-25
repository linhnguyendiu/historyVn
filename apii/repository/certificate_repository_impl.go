package repository

import (
	"errors"
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/model/domain"

	"gorm.io/gorm"
)

type CertificateRepositoryImpl struct {
	db *gorm.DB
}

func (r *CertificateRepositoryImpl) Save(certificate domain.Certificate) domain.Certificate {
	err := r.db.Create(&certificate).Error
	helper.PanicIfError(err)

	return certificate
}

func (r *CertificateRepositoryImpl) FindById(certId int) (domain.Certificate, error) {
	certificate := domain.Certificate{}
	err := r.db.Find(&certificate, "id=?", certId).Error
	if certificate.Id == 0 || err != nil {
		return certificate, errors.New("certificate not found")
	}

	return certificate, nil
}

func NewCertificateRepository(db *gorm.DB) CertificateRepository {
	return &CertificateRepositoryImpl{db: db}
}
