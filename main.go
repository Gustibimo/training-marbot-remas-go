package main

import (
	"errors"
	"fmt"
	"net/http"
	"github.com/Gustibimo/training-marbot-remas-go.git/presentation"
)

func main() {

	http.HandleFunc("/", presentation.Root)
	http.HandleFunc("/users/", presentation.GetUserByFirstName)

	err := http.ListenAndServe(":3333", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed")
	} else if err != nil {
		fmt.Printf("error starting server")
	}

}
