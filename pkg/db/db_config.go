package db

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	mod "github.com/shivaraj-shanthaiah/godoc-user/pkg/models"
	
)

func ConnectDB() *gorm.DB{
	dsn, ok := os.LookupEnv("DB_Config")
	if !ok {
		log.Fatal("Error handling database env")
	}

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err !=nil{
		log.Fatalf("Error connecting to databse: %v", err.Error())
	}
	DB.AutoMigrate(&mod.User{})
	return DB
}