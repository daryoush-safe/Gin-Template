package application_math

import (
	"gingool/src/repository"
)

type AddService struct {
	userRepository *repository.UserRepository
}

func NewAddService(userRepository *repository.UserRepository) *AddService {
	return &AddService{
		userRepository: userRepository,
	}
}

func (addService *AddService) Add(num1 int, num2 int) int {
	return num1 + num2
}
