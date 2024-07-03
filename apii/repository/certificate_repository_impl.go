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

func (r *CertificateRepositoryImpl) FindById(certId int, userId int) (domain.Certificate, error) {
	certificate := domain.Certificate{}
	err := r.db.First(&certificate, "id = ? AND user_id = ?", certId, userId).Error
	if err != nil {
		return certificate, errors.New("certificate not found")
	}

	return certificate, nil
}

func NewCertificateRepository(db *gorm.DB) CertificateRepository {
	return &CertificateRepositoryImpl{db: db}
}
