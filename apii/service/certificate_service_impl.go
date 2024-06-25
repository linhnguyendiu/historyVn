package service

import (
	"bytes"
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/model/domain"
	"go-pzn-restful-api/model/web"
	"go-pzn-restful-api/repository"
	"log"

	"github.com/jung-kurt/gofpdf"
	"github.com/jung-kurt/gofpdf/contrib/httpimg"
)

type CertificateServiceImpl struct {
	repository.CertificateRepository
}

func (s *CertificateServiceImpl) Create(input web.CertificateCreateInput) web.CertificateResponse {

	cert := domain.Certificate{
		UserName:   input.UserName,
		CourseName: input.CourseName,
		Date:       input.Date,
		CertType:   input.CertType,
		ImageUri:   input.ImageUri,
	}

	certificatePDF, err := GenerateCertificatePDF(cert)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	// srv, err := helper.NewDriveService()
	// if err != nil {
	// 	panic(helper.NewNotFoundError(err.Error()))
	// }

	size := int64(len(certificatePDF))

	driveFileID, err := helper.CreateFile(cert.CourseName+".pdf", size, certificatePDF)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	cert.CertUri = driveFileID
	certificateNFT := s.CertificateRepository.Save(cert)

	return helper.ToCertificateResponse(certificateNFT)
}

func GenerateCertificatePDF(req domain.Certificate) ([]byte, error) {
	// Tạo file PDF
	pdf := gofpdf.New("P", "mm", "A6", "")
	pdf.AddPage()

	// // Thêm logo
	// pdf.Image("../assets/c", 10, 10, 30, 0, false, "", 0, "")

	// Thêm tiêu đề
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "CERTIFICATE OF AUTHENTICITY")

	// Thêm hình ảnh minh họa
	httpimg.Register(pdf, req.ImageUri, "")
	if pdf.Err() {
		log.Printf("error registering image %s: %s", req.ImageUri, pdf.Error())
		pdf.ClearError()
	}
	pdf.Image(req.ImageUri, 10, 30, 80, 0, false, "", 0, "")

	// Thêm thông tin người học và khóa học
	pdf.SetXY(100, 30)
	pdf.SetFont("Arial", "B", 14)
	pdf.MultiCell(100, 10, req.UserName, "", "L", false)
	pdf.MultiCell(100, 10, req.CourseName, "", "L", false)
	// pdf.MultiCell(100, 10, req.Date, "", "L", false)

	// // Thêm chữ ký
	// pdf.Image("path/to/signature.png", 100, 80, 30, 0, false, "", 0, "")
	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (s *CertificateServiceImpl) FindById(certId int) web.CertificateResponse {
	findById, err := s.CertificateRepository.FindById(certId)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	return helper.ToCertificateResponse(findById)
	//return helper.ToCourseResponse(findById, countUsersEnrolled)
}

// func (s *CertificateServiceImpl) DownloadCertificateFile(certId int) ([]byte, error) {
// 	findById, err := s.CertificateRepository.FindById(certId)
// 	if err != nil {
// 		panic(helper.NewNotFoundError(err.Error()))
// 	}
// 	file, err := s.Drive.Files.Get(findById.CertUri).Download()
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer file.Body.Close()

// 	data, err := io.ReadAll(file.Body)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return data, nil
// }

func NewCertificateService(certificateRepository repository.CertificateRepository) CertificateService {
	return &CertificateServiceImpl{CertificateRepository: certificateRepository}
}
