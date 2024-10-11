package application

import (
	"first-project/src/bootstrap"
	"first-project/src/exceptions"
	"first-project/src/repository"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	constants      *bootstrap.Constants
	userRepository *repository.UserRepository
}

func NewUserService(constants *bootstrap.Constants, userRepository *repository.UserRepository) *UserService {
	return &UserService{
		constants:      constants,
		userRepository: userRepository,
	}
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func validatePasswordTests(errors *[]string, test string, password string, tag string) {
	matched, _ := regexp.MatchString(test, password)
	if !matched {
		*errors = append(*errors, tag)
	}
}

func (userService *UserService) passwordValidation(password string) []string {
	var errors []string
	validatePasswordTests(&errors, ".{8,}", password, userService.constants.Context.MinimumLength)
	validatePasswordTests(&errors, "[a-z]", password, userService.constants.Context.ContainsLowercase)
	validatePasswordTests(&errors, "[A-Z]", password, userService.constants.Context.ContainsUppercase)
	validatePasswordTests(&errors, "[0-9]", password, userService.constants.Context.ContainsNumber)
	validatePasswordTests(&errors, "[^\\d\\w]", password, userService.constants.Context.ContainsSpecialChar)

	return errors
}

func (userService *UserService) RegisterService(username string, email string, password string) {
	var registrationError exceptions.UserRegistrationError
	isRegError := false
	usernameExist := userService.userRepository.CheckUsernameExists(username)
	if usernameExist {
		isRegError = true
		// TODO: why context?
		// TODO: use translate here ?! not sure =)
		registrationError.AppendError("Username", userService.constants.Context.AlreadyExist)
	}
	emailExist := userService.userRepository.CheckEmailExists(email)
	if emailExist {
		isRegError = true
		registrationError.AppendError("Email", userService.constants.Context.AlreadyExist)
	}
	passwordErrorTags := userService.passwordValidation(password)
	if len(passwordErrorTags) > 0 {
		isRegError = true
		for _, v := range passwordErrorTags {
			registrationError.AppendError("Password", v)
		}
	}

	if isRegError {
		panic(registrationError)
	}
	hashedPassword, err := hashPassword(password)
	if err != nil {
		panic(err)
	}
	userService.userRepository.RegisterUser(username, email, hashedPassword)
}

func (userService *UserService) VerifyEmail(email string) {
	var registrationError exceptions.UserRegistrationError
	alreadyVerified := userService.userRepository.CheckUserVerified(email)
	if alreadyVerified {
		registrationError.AppendError("Email", userService.constants.Context.AlreadyVerified)
		panic(registrationError)
	}
	userService.userRepository.VerifyEmail(email)
}
