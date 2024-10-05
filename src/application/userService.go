package application

import (
	"first-project/src/exceptions"
	"first-project/src/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewRegisterService(userRepository *repository.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (userService *UserService) RegisterService(username string, email string, password string) {
	var regError exceptions.UserRegistrationError // returned Error func!! issue on both of errors
	isRegError := false
	usernameExist := userService.userRepository.CheckUsernameExists(username)
	if usernameExist {
		// usernameExistError := exceptions.UserRegistrationError{Username: username}
		// panic(usernameExistError)
		isRegError = true
		regError.Username = username
	}
	emailExist := userService.userRepository.CheckEmailExists(email)
	if emailExist {
		// emailExistError := exceptions.UserRegistrationError{Email: email}
		// panic(emailExistError)
		isRegError = true
		regError.Email = email
	}
	if isRegError {
		panic(regError)
	}
	hashedPassword, err := hashPassword(password)
	if err != nil {
		panic(err)
	}
	userService.userRepository.RegisterUser(username, email, hashedPassword)
}
