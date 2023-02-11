package persistence

import (
	"fmt"
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

func (u *Users) FindFirstName() {
	path := "./db/users.csv"
	lines := make(chan string)

	go readFile(path, lines)

	for line := range lines {
		user := parseLine(line)
		u = &user
		fmt.Println(u.first_name)
	}
}

func (u *Users) FindUserByFirstName(first string) *Users {
	path := "./db/users.csv"
	lines := make(chan string)

	go readFile(path, lines)

	for line := range lines {
		user := parseLine(line)
		u = &user
		if u.first_name == first {
			fmt.Println(u)
		}
	}
	fmt.Println("Not found")
	return nil
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
