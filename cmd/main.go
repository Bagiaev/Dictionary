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

	router := echo.New()
	// создаем группу api
	api := router.Group("api")

	// прописываем пути
	api.GET("/word/:id", svc.GetWordById)
	api.POST("/words", svc.CreateWords)

	//Новые ручки "Задание 1"
	api.PUT("/word/:id", svc.UpdateWord)    // Обновление
	api.DELETE("/word/:id", svc.DeleteWord) // Удаление

	//Задание 2
	//Пути reports
	api.GET("/report/:id", svc.GetReport)
	api.PATCH("/report/:id", svc.UpdateReport)
	api.POST("/report", svc.CreateReport)
	api.DELETE("/report/:id", svc.DeleteReport)

	// запускаем сервер, чтобы слушал 8000 порт
	router.Logger.Fatal(router.Start(":8000"))
}
