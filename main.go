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
	r.Use(corsMiddleware())

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

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Method", "*")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	}
}
