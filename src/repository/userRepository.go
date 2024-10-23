package repository

import (
	"first-project/src/entities"

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

func (repo *UserRepository) FindByUsernameAndVerified(username string, verified bool) (entities.User, bool) {
	var user entities.User
	result := repo.db.Where("name = ? AND verified = ?", username, verified).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return user, false
		}
		panic(result.Error)
	}
	return user, true
}

func (repo *UserRepository) FindByEmailAndVerified(email string, verified bool) (entities.User, bool) {
	var user entities.User
	result := repo.db.Where("email = ? AND verified = ?", email, verified).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return user, false
		}
		panic(result.Error)
	}
	return user, true
}

func (repo *UserRepository) UpdateUserToken(user entities.User, token string) {
	user.Token = token
	repo.db.Save(&user)
}

func (repo *UserRepository) CreateNewUser(username string, email string, password string, token string, verified bool) {
	user := entities.User{
		Name:     username,
		Email:    email,
		Password: password,
		Token:    token,
		Verified: verified,
	}
	result := repo.db.Create(&user)
	if result.Error != nil {
		panic(result.Error)
	}
}

func (repo *UserRepository) ActivateUserAccount(user entities.User) {
	user.Verified = true
	user.Token = ""
	if err := repo.db.Save(&user).Error; err != nil {
		panic(err)
	}
}

func (repo *UserRepository) UpdateUserPassword(user entities.User, password string) {
	user.Password = password
	user.Token = ""
	repo.db.Save(&user)
}
