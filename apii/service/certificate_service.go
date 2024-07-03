package service

import "go-pzn-restful-api/model/web"

type CertificateService interface {
	Create(input web.CertificateCreateInput) web.CertificateResponse
	FindById(certId int, userId int) web.CertificateResponse
}
