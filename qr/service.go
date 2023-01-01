package qr

import (
	"errors"

	"github.com/skip2/go-qrcode"
)

// interface
type IQRService interface {
	GenerateQR(input string) ([]byte, error)
}

type QRService struct{}

// new service
func NewQRService() *QRService {
	return &QRService{}
}

// implementation
func (s *QRService) GenerateQR(input string) ([]byte, error) {
	if input == "" {
		return nil, errors.New("input cannot be empty")
	}

	qrCode, err := qrcode.Encode(input, qrcode.Medium, 256)
	if err != nil {
		return nil, err
	}

	return qrCode, nil
}
