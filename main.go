package main

import "github.com/Gustibimo/training-marbot-remas-go.git/persistence"

func main() {
	var userRepo persistence.Users
	userRepo.FindUserByFirstName("Rebecca")
}
