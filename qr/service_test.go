package qr_test

import (
	"testing"

	"github.com/rezairfanwijaya/go-qr-backend/qr"
	"github.com/stretchr/testify/assert"
)

func TestQRGenerate(t *testing.T) {
	testCasaes := []struct {
		Name      string
		Input     string
		WantError bool
	}{
		{
			Name:      "success",
			Input:     "github.com/rezairfanwijaya/go-qr-backend/qr",
			WantError: false,
		}, {
			Name:      "failed",
			Input:     "",
			WantError: true,
		},
	}

	for _, testCase := range testCasaes {
		t.Run(testCase.Name, func(t *testing.T) {
			QRService := qr.QRService{}

			actual, err := QRService.GenerateQR(testCase.Input)
			if testCase.WantError {
				assert.NotNil(t, err)
				assert.Nil(t, actual)
			} else {
				assert.NotNil(t, actual)
				assert.Nil(t, err)
			}
		})
	}

}
