package repositories

import (
	"github.com/IgorDePaula/plataforma-go/domain"
	"github.com/jinzhu/gorm"
	"log"
)

type UserRepository interface {
	Insert(user *domain.User) (*domain.User, error)
}

type UserRepositoryDb struct {
	Db *gorm.DB
}

func (repo UserRepositoryDb) Insert (user *domain.User) (*domain.User, error){
	err := user.Prepare()

	if err != nil{
		log.Fatalf("errro validate user %v", err)
	}

	err = repo.Db.Create(user).Error

	if err != nil{
		log.Fatalf("errro persist user %v", err)
		return user, err
	}
	return user, nil
}