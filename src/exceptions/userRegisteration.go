package exceptions

import (
	"fmt"
)

// type UsernameExistsError struct {
// 	Username string
// }

// func (e UsernameExistsError) Error() string {
// 	return fmt.Sprintf("Username %s already exists.", e.Username)
// }

// type EmailExistsError struct {
// 	Email string
// }

// func (e EmailExistsError) Error() string {
// 	return fmt.Sprintf("Email %s already exists.", e.Email)
// }

type UserRegistrationError struct {
	Username string
	Email    string
}

func (e UserRegistrationError) Error() string {
	// var massages []string
	// if e.Username != "" {
	// 	massages = append(massages, fmt.Sprintf("Username %s already exists.", e.Username))
	// } else if e.Email != "" {
	// 	massages = append(massages, fmt.Sprintf("Email %s already exists.", e.Email))
	// } else {
	// 	return "Registration failed."
	// }
	// return strings.Join(massages, "; ")
	if e.Username != "" {
		return fmt.Sprintf("Username %s already exists.", e.Username)
	} else if e.Email != "" {
		return fmt.Sprintf("Email %s already exists.", e.Email)
	}
	return "Registration failed."
}
