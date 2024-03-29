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
	"intern-bcc/pkg/midtranss"
	"intern-bcc/pkg/supabase"
)

func main() {

	config.LoadEnv()

	jwtAuth := jwt.Init()

	bCrypt := bcrypt.Init()

	supabase := supabase.Init()

	db := mysql.ConnectDatabase()

	newRepository := repository.NewRepository(db)

	newTopUp := repository.NewTopUp(db)

	midtrans := midtranss.NewMidtrans(&midtranss.Midtrans{})

	newService := service.NewService(service.InitParam{
		Repository: newRepository,
		JwtAuth: jwtAuth,
		Bcrypt: bCrypt,
		Supabase: supabase,
		TopUp: newTopUp,
		Midtrans: midtrans,
	})

	middleware := middleware.Init(jwtAuth,newService)

	newHandler := handler.NewHandler(newService, middleware)

	mysql.Migration(db)

	newHandler.EndPoint()
}
