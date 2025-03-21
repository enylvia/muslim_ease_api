package main

import (
	"context"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go-api-quran/app"
	_ "go-api-quran/docs"
	"go-api-quran/handler"
	"go-api-quran/repository"
	"go-api-quran/service"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// @title MuslimEase - Developer API
// @version 1.0
// @description A modern and open Islamic API providing Qur'an data, prayer schedules, daily duas, and more — built for developers to create powerful and meaningful Islamic applications.
// @host localhost:8080
// @BasePath /api/v1
func main() {
	dbInitialize := app.Init()
	surahRepository := repository.NewSurahRepository(dbInitialize.DB)
	surahService := service.NewSurahService(surahRepository)
	surahHandler := handler.NewSurahHandler(surahService)

	setupLogging()
	r := gin.Default()
	r.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
	}))

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api/")
	v1 := api.Group("/v1/")
	{
		surah := v1.Group("/surah/")
		{
			surah.GET("", surahHandler.ListSurah)
		}
	}
	startWithGracefulShutdown(r)
}

func setupLogging() {
	fLog, err := os.Create("logs/gin.log")
	if err != nil {
		log.Fatalf("Failed to create log file: %v", err)
	}
	gin.DefaultWriter = io.MultiWriter(os.Stdout, fLog)
}

func startWithGracefulShutdown(router *gin.Engine) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	go func() {
		log.Printf("🚀 Server running at port: %s", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %s", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("📴 Shutting down server...")

	// Timeout shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("❌ Server forced to shutdown: %s", err)
	}

	log.Println("✅ Server exited gracefully")
}
