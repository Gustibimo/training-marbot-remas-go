package main

import (
	"errors"
	"fmt"
	"net/http"

	// "github.com/Gustibimo/training-marbot-remas-go.git/domain"
	// "github.com/Gustibimo/training-marbot-remas-go.git/persistence"
	"github.com/Gustibimo/training-marbot-remas-go.git/presentation"
)

func main() {
	// var userRepo persistence.Users
	// userRepo.FindUserByFirstName("Rebecca")

	// var user persistence.Users

	// userService := domain.NewUserService(&user)

	// result, err := userService.FindUserByFirstName("Rebecca")
	// if err != nil {
	// 	fmt.Println("not found")
	// }
	// fmt.Print(result)

	// res := userService.FindAllUserFirstName()
	// fmt.Println(res)

	http.HandleFunc("/", presentation.Root)
	http.HandleFunc("/users/", presentation.GetUserByFirstName)

	err := http.ListenAndServe(":3333", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed")
	} else if err != nil {
		fmt.Printf("error starting server")
	}
	


}
