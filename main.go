package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "go-api-quran/docs"
	"io"
	"log"
	"os"
)

// @title Contoh API dengan Swagger
// @version 1.0
// @description Ini adalah contoh dokumentasi Swagger dengan Gin
// @host localhost:8080
// @BasePath /
func main() {
	//dbInitialize := app.Init()
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/hello", helloHandler)

	r.Run(":8080")
}

func setupLogging() {
	fLog, err := os.Create("logs/gin.log")
	if err != nil {
		log.Fatalf("Failed to create log file: %v", err)
	}
	gin.DefaultWriter = io.MultiWriter(os.Stdout, fLog)
}

// @Summary Get Pong
// @Description Mengembalikan respon "pong"
// @Tags Example
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]string
// @Router /hello [get]
func helloHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Hello World"})
}
