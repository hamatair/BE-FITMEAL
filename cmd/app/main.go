package main

import (
	"intern-bcc/internal/handler"
	"intern-bcc/internal/repository"
	"intern-bcc/internal/service"
	"intern-bcc/pkg/config"
	"intern-bcc/pkg/database/mysql"
)

func main() {
	config.LoadEnv()

	db := mysql.ConnectDatabase()

	newRepository := repository.NewRepository(db)

	newService := service.NewService(newRepository)

	newHandler := handler.NewHandler(newService)

	mysql.Migration(db)

	handler.EndPoint(newHandler)
}
