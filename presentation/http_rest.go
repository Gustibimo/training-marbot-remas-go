package presentation

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Gustibimo/training-marbot-remas-go.git/domain"
	"github.com/Gustibimo/training-marbot-remas-go.git/persistence"
)

func Root(w http.ResponseWriter, r *http.Request)  {

	response := fmt.Sprintf("Simple server is up %v", http.StatusOK)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

func GetUserByFirstName(w http.ResponseWriter, r *http.Request) {

	firstname := r.URL.Query().Get("first_name")
	
	var user persistence.Users

	userService := domain.NewUserService(&user)

	result, err := userService.FindUserByFirstName(firstname)
	if err != nil {
		fmt.Println("not found")
		log.Fatalf("error not found")
	}

	fmt.Println(result)

	jsonstring, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	io.WriteString(w, string(jsonstring))

}


