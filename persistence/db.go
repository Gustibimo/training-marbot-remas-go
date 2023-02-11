package persistence

import (
	"bufio"
	"log"
	"os"
)

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