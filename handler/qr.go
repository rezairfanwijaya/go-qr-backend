package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rezairfanwijaya/go-qr-backend/qr"
	"github.com/rezairfanwijaya/go-qr-backend/utils"
)

type QRHandler struct {
	qrService qr.IQRService
}

// new handler
func NewQRHandler(qrService qr.IQRService) *QRHandler {
	return &QRHandler{qrService}
}

// implementation
func (h *QRHandler) GenerateQR(c *gin.Context) {
	var input qr.InputBodyQR

	// binding
	if err := c.ShouldBindJSON(&input); err != nil {
		response := utils.ResponseAPI(
			"failed",
			err.Error(),
			http.StatusBadRequest,
			nil,
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	// call service
	qrCode, err := h.qrService.GenerateQR(input.Input)
	if err != nil {
		response := utils.ResponseAPI(
			"failed",
			err.Error(),
			http.StatusBadRequest,
			nil,
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.ResponseAPI(
		"success",
		"success to generate QR Code",
		http.StatusOK,
		qrCode,
	)
	c.JSON(http.StatusOK, response)
}
