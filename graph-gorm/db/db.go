package db

import (
	"log"

	"github.com/99designs/gqlgen/codegen/testserver/benchmark/generated/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=sanufinu786 dbname=graphql_db port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn),
		&gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect db")
	}
	db.AutoMigrate(&models.User{})

	DB = db

}
