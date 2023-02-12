package persistence

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	db "github.com/Gustibimo/training-marbot-remas-go.git/db"
)

type Users struct {
	Id             int
	First_name     string
	Last_name      string
	Email_address  string
	Created_at     int
	Deleted_at     int
	Merged_at      int
	Parent_user_id int
}

type UserResponse struct {
	Id           int    `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	EmailAddress string `json:"email_address"`
	CreatedAt    int    `json:"created_at"`
	DeletedAt    int    `json:"deleted_at"`
	MergedAt     int    `json:"merged_at"`
	ParentUserId int    `json:"parent_user_id"`
}

// overrride MarshalJSON from Marshaler interface
func (u Users) MarshalJSON() ([]byte, error) {
	return json.Marshal(UserResponse{
		Id:           u.Id,
		FirstName:    u.First_name,
		LastName:     u.Last_name,
		EmailAddress: u.Email_address,
		CreatedAt:    u.Created_at,
		DeletedAt:    u.Deleted_at,
		MergedAt:     u.Merged_at,
		ParentUserId: u.Parent_user_id,
	})
}

func (u *Users) FindFirstName() []string {
	path := "./db/users.csv"
	lines := make(chan string)

	var users []string

	go db.ReadFile(path, lines)

	for line := range lines {
		user := parseLine(line)
		u = &user
		users = append(users, u.First_name)
	}

	return users
}

func (u *Users) FindUserByFirstName(first string) ([]Users, error) {
	path := "./db/users.csv"
	lines := make(chan string)

	var users []Users

	go db.ReadFile(path, lines)

	for line := range lines {
		user := parseLine(line)
		u = &user
		if u.First_name == first {
			users = append(users, user)
		}
	}
	if len(users) == 0 {
		return nil, errors.New("not found")
	}
	return users, nil
}

func parseLine(line string) Users {

	values := strings.Split(line, ",")
	id, _ := strconv.Atoi(values[0])
	created_at, _ := strconv.Atoi(values[4])
	deleted_at, _ := strconv.Atoi(values[5])
	merged_at, _ := strconv.Atoi(values[6])
	parent_user_id, _ := strconv.Atoi(values[7])

	return Users{
		Id:             id,
		First_name:     values[1],
		Last_name:      values[2],
		Email_address:  values[3],
		Created_at:     created_at,
		Deleted_at:     deleted_at,
		Merged_at:      merged_at,
		Parent_user_id: parent_user_id,
	}

}
