package main

import (
	"fmt"

	"github.com/Gustibimo/training-marbot-remas-go.git/domain"
	"github.com/Gustibimo/training-marbot-remas-go.git/persistence"
)

func main() {
	// var userRepo persistence.Users
	// userRepo.FindUserByFirstName("Rebecca")

	var user persistence.Users

	userService := domain.NewUserService(&user)

	result := userService.FindUserByFirstName("Rebecca")
	fmt.Print(result)

	// res := userService.FindAllUserFirstName()
	// fmt.Println(res)


}
