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

	result, err := userService.FindUserByFirstName("Rebecca")
	if err != nil {
		fmt.Println("not found")
	}
	fmt.Print(result)

	// res := userService.FindAllUserFirstName()
	// fmt.Println(res)


}
