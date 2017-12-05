package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	validCount := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if isValid(line) {
			validCount++
		} else {
			fmt.Println("Not Valid passphrase: ", line)
		}
	}

	fmt.Printf("Valid passphrases count: %d", validCount)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func isValid(passphrase string) bool {
	var parts = strings.Split(passphrase, " ")

	var m = make(map[string]bool)

	for i := 0; i < len(parts); i++ {
		var p = parts[i]

		_, found := m[p]

		if found {
			return false
		}
		m[p] = true
	}

	return true
}
