package service

import (
	"bytes"
	"errors"
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/model/domain"
	"go-pzn-restful-api/model/web"
	"go-pzn-restful-api/repository"
	"image"
	"net/http"

	"github.com/fogleman/gg"
	"github.com/nfnt/resize"
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

	certificatePDF, err := GenerateCertNFTPDF(cert)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	// srv, err := helper.NewDriveService()
	// if err != nil {
	// 	panic(helper.NewNotFoundError(err.Error()))
	// }

	size := int64(len(certificatePDF))

	driveFileID, err := helper.CreateFile(cert.CourseName+".png", size, certificatePDF)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	cert.CertUri = driveFileID
	certificateNFT := s.CertificateRepository.Save(cert)

	return helper.ToCertificateResponse(certificateNFT)
}

func GenerateCertNFTPDF(cert domain.Certificate) ([]byte, error) {
	// Tạo khung ảnh với kích thước cố định
	const widthPx, heightPx = 1200, 800
	dc := gg.NewContext(widthPx, heightPx)

	// Background màu trắng
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	// Vẽ tiêu đề chứng chỉ
	dc.SetRGB(0, 0, 0)
	if err := dc.LoadFontFace("assets/font/BeVietnamPro-Black.ttf", 36); err != nil {
		return nil, err
	}
	dc.DrawStringAnchored("CHỨNG NHẬN HOÀN THÀNH KHÓA HỌC", widthPx/2, 100, 0.5, 0.5)

	// Vẽ hình ảnh chính với kích thước cố định
	response, err := http.Get(cert.ImageUri)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, errors.New("received non 200 response code")
	}

	img, _, err := image.Decode(response.Body)
	if err != nil {
		return nil, err
	}

	imgWidth, imgHeight := 300, 450
	resizedImg := resize.Resize(uint(imgWidth), uint(imgHeight), img, resize.Lanczos3)
	dc.DrawImageAnchored(resizedImg, 200, 400, 0.5, 0.5)

	// Vẽ viền (border) cho phần text
	borderX := 550.0
	borderY := 150.0
	borderHeight := 500.0
	borderThickness := 2.0

	dc.SetLineWidth(borderThickness)
	dc.SetRGB(0, 0, 0)
	dc.DrawLine(borderX, borderY, borderX, borderY+borderHeight)
	dc.Stroke()

	// Vẽ thông tin chi tiết chứng chỉ
	if err := dc.LoadFontFace("assets/font/BeVietnamPro-Black.ttf", 24); err != nil {
		return nil, err
	}
	startX := 600.0
	startY := 200.0
	lineHeight := 40.0

	dc.DrawStringAnchored("Người nhận:", startX, startY, 0, 0.5)
	dc.DrawStringAnchored(cert.UserName, startX, startY+lineHeight, 0, 0.5)
	dc.DrawStringAnchored("Khóa học:", startX, startY+2*lineHeight, 0, 0.5)
	dc.DrawStringAnchored(cert.CourseName, startX, startY+3*lineHeight, 0, 0.5)
	dc.DrawStringAnchored("Ngày hoàn thành:", startX, startY+4*lineHeight, 0, 0.5)
	dc.DrawStringAnchored(cert.Date.Format("02-01-2006"), startX, startY+5*lineHeight, 0, 0.5)
	dc.DrawStringAnchored("Loại chứng chỉ:", startX, startY+6*lineHeight, 0, 0.5)
	dc.DrawStringAnchored(cert.CertType, startX, startY+7*lineHeight, 0, 0.5)

	var buf bytes.Buffer
	err = dc.EncodePNG(&buf)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (s *CertificateServiceImpl) FindById(certId int, userId int) web.CertificateResponse {
	findById, err := s.CertificateRepository.FindById(certId, userId)
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
