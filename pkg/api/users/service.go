package users

import (
	"fmt"
)

// User the model to store the user information
type User struct {
	ID       int64  `json:"id,omitempty"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

type userService struct {
	repo repo
}

func newUserService() *userService {
	return &userService{repo: newInMemoryStore()}
}

func (us *userService) RegisterUser(user User) (User, error) {
	u, err := us.repo.CreateUser(user)
	if err != nil {
		fmt.Println("Error... while handling tools.")
	}
	return u, nil
}
