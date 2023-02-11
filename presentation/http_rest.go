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

type UserResponse struct {
	Id             int  `json:"id" bson:"id"`
	FirstName     string `json:"firstName" bson:"first_name"`
	LastName      string `json:"lastName" bson:"last_name"`
	EmailAddress  string `json:"emailAddress" bson:"email_address"`
	CreatedAt     int `json:"createdAt" bson:"created_at"`
	DeletedAt     int `json:"deletedAt" bson:"deleted_at"`
	MergedAt      int `json:"mergedAt" bson:"merged_at"`
	ParentUserId int `json:"parentUserId" bson:"parent_user_id"`
}

func Root(w http.ResponseWriter, r *http.Request)  {

	firstname := r.URL.Query().Get("first_name")

	response := fmt.Sprintf("hello %s", firstname)

	io.WriteString(w, response)
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


