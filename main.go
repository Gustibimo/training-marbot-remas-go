package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func main() {

	path := "./db/users.csv"
	lines := make(chan string)

	var user Users

	go readFile(path, lines)

	for line := range lines {
		user = parseLine(line)
		user.PrintFirst()
	}
}

func readFile(path string, lines chan<- string) {
	file, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines <- scanner.Text()
	}

	close(lines)
}

func (u *Users) PrintFirst() {
	fmt.Println(u.first_name)
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
