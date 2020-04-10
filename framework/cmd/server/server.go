package main

import (
	"fmt"
	"github.com/IgorDePaula/plataforma-go/application/repositories"
	"github.com/IgorDePaula/plataforma-go/domain"
	"github.com/IgorDePaula/plataforma-go/framework/utils"
	"log"
)

func main() {

	db := utils.ConnectDB()

	user := domain.User{
		Name:     "Igor",
		Email:    "igor@example.org",
		Password: "123",
	}

	userRepo := repositories.UserRepositoryDb{Db: db}
	result, err := userRepo.Insert(&user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.Name, result.Email, result.Token)

}
