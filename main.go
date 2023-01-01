package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rezairfanwijaya/go-qr-backend/handler"
	"github.com/rezairfanwijaya/go-qr-backend/qr"
)

func main() {
	// init gin
	r := gin.Default()

	// init service qr
	service := qr.NewQRService()
	// init handler qr
	handler := handler.NewQRHandler(service)

	// endpoint
	r.POST("generate", handler.GenerateQR)

	// run server
	if err := r.Run(":8989"); err != nil {
		log.Fatal(err)
	}
}
