package main

import (
	"intern-bcc/internal/handler"
	"intern-bcc/internal/repository"
	"intern-bcc/internal/service"
	"intern-bcc/pkg/bcrypt"
	"intern-bcc/pkg/config"
	"intern-bcc/pkg/database/mysql"
	"intern-bcc/pkg/jwt"
	"intern-bcc/pkg/middleware"
)

func main() {
	config.LoadEnv()

	jwtAuth := jwt.Init()

	bCrypt := bcrypt.Init()

	db := mysql.ConnectDatabase()

	newRepository := repository.NewRepository(db)

	newService := service.NewService(service.InitParam{Repository: newRepository, JwtAuth: jwtAuth, Bcrypt: bCrypt})

	middleware := middleware.Init(jwtAuth,newService)

	newHandler := handler.NewHandler(newService, middleware)

	mysql.Migration(db)

	newHandler.EndPoint()
}
