package domain

import "github.com/Gustibimo/training-marbot-remas-go.git/persistence"


type UserFinder interface {
	FindUserByFirstName(firstName string) []persistence.Users
	FindAllUserFirstName() []string
}

type UserWriter interface {
	CreateUser(user persistence.Users) error
	UpdateUser(user persistence.Users) error
	DeleteUser(user persistence.Users) error
}

type UserService struct {
	Users *persistence.Users
}

func NewUserService(users *persistence.Users) UserFinder {
	return &UserService{
		Users: users,
	}
}

func (s *UserService) FindUserByFirstName(firstName string) []persistence.Users {
	return s.Users.FindUserByFirstName(firstName)
}

func (s *UserService) FindAllUserFirstName() []string {
	// Implementation to return all users from the persistence layer
	return s.Users.FindFirstName()
}

func (s *UserService) CreateUser(user persistence.Users) error {
	// Implementation to create a new user
	return nil
}

func (s *UserService) UpdateUser(user persistence.Users) error {
	// Implementation to update an existing user
	return nil
}

func (s *UserService) DeleteUser(user persistence.Users) error {
	// Implementation to delete a user
	return nil
}

