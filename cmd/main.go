package main

import (
	"dictionary/internal/service"
	"dictionary/pkg/logs"

	"github.com/labstack/echo/v4"
)

func main() {
	// создаем логгер
	logger := logs.NewLogger(false)

	// подключаемся к базе
	db, err := PostgresConnection()
	if err != nil {
		logger.Fatal(err)
	}

	svc := service.NewService(db, logger)
	reportsSvc := service.NewReportsService(db, logger)

	router := echo.New()
	// создаем группу api
	api := router.Group("api")

	// прописываем пути
	api.GET("/word/:id", svc.GetWordById)
	api.POST("/words", svc.CreateWords)
	api.PUT("/word/:id", svc.UpdateWord)
	api.DELETE("/word/:id", svc.DeleteWord)

	// пути для репортов
	api.GET("/report/:id", reportsSvc.GetReportById)
	api.POST("/report", reportsSvc.CreateReport)
	api.PUT("/report/:id", reportsSvc.UpdateReport)
	api.DELETE("/report/:id", reportsSvc.DeleteReport)

	// запускаем сервер, чтобы слушал 8000 порт
	router.Logger.Fatal(router.Start(":8080"))
}
