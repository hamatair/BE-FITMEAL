package mysql

import (
	"intern-bcc/pkg/config"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {

	db, err := gorm.Open(mysql.Open(config.LoadDatabaseConfig()), &gorm.Config{})
	if err != nil {
		log.Fatal("Error to connecting database")
	}

	return db
}