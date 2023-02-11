package persistence

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
)

type Users struct {
	id             int
	first_name     string
	last_name      string
	email_address  string
	created_at     int
	deleted_at     int
	merged_at      int
	parent_user_id int
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

func (u Users) MarshalJSON() ([]byte, error) {
	return json.Marshal(UserResponse{
		Id:           u.id,
		FirstName:    u.first_name,
		LastName:     u.last_name,
		EmailAddress: u.email_address,
		CreatedAt:    u.created_at,
		DeletedAt:    u.deleted_at,
		MergedAt:     u.merged_at,
		ParentUserId: u.parent_user_id,
	})
}

func (u *Users) FindFirstName() []string {
	path := "./db/users.csv"
	lines := make(chan string)

	var users []string

	go readFile(path, lines)

	for line := range lines {
		user := parseLine(line)
		u = &user
		users = append(users, u.first_name)
	}

	return users
}

func (u *Users) FindUserByFirstName(first string) ([]Users, error) {
	path := "./db/users.csv"
	lines := make(chan string)

	var users []Users

	go readFile(path, lines)

	for line := range lines {
		user := parseLine(line)
		u = &user
		if u.first_name == first {
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
		id:             id,
		first_name:     values[1],
		last_name:      values[2],
		email_address:  values[3],
		created_at:     created_at,
		deleted_at:     deleted_at,
		merged_at:      merged_at,
		parent_user_id: parent_user_id,
	}

}
