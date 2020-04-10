package utils

import (
	"github.com/IgorDePaula/plataforma-go/domain"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"log"
	"os"
	_ "github.com/lib/pq"
)

func ConnectDB() *gorm.DB{

	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error load env %v", err)
	}

	dsn := os.Getenv("dsn")

	db, err := gorm.Open("postgres", dsn)

	// defer db.Close()

	if err != nil{
		log.Fatalf("Error connect in database %v", err)
		panic(err)
	}

	db.AutoMigrate(&domain.User{})

	return db
}
