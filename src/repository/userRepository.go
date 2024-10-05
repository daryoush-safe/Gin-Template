package repository

import (
	"first-project/src/entities"
	"log"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (repo *UserRepository) Test() []string {
	var tables []string
	repo.db.Raw("SHOW TABLES").Scan(&tables)

	return tables
}

func (repo *UserRepository) Test2() []entities.Test {
	var results []entities.Test
	repo.db.Where("name = ?", "ali").Find(&results)

	return results
}

func (repo *UserRepository) CheckUsernameExists(username string) bool {
	var user entities.User
	result := repo.db.Where("name = ?", username).First(&user)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return false
		}
		log.Println("Error occurred:", result.Error)
		return true
	}

	return true
}

func (repo *UserRepository) CheckEmailExists(email string) bool {
	var user entities.User
	result := repo.db.Where("email = ?", email).First(&user)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return false
		}
		log.Println("Error occurred:", result.Error)
		return true
	}

	return true
}

func (repo *UserRepository) RegisterUser(username string, email string, password string) {
	user := entities.User{
		Name:     username,
		Email:    email,
		Password: password,
	}
	result := repo.db.Create(&user)
	if result.Error != nil {
		log.Println("Error creating user: ", result.Error)
	}
}
