package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type Group struct {
}

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	fileLines, err := read(file)

	line := fileLines[0]

	score := parse(line)

	fmt.Printf("Score %d\n", score)
}

func parse(input string) int {
	score := 0
	group := 0
	garbage := false
	cancel := false

	for _, char := range input {
		if cancel {
			cancel = false
			continue
		}
		if char == '!' && cancel == false {
			cancel = true
		}

		if cancel == false {
			if char == '<' {
				garbage = true
				continue
			}
			if char == '>' {
				garbage = false
				continue
			}
			if garbage == false && char == '{' {
				group++
			}
			if garbage == false && char == '}' {
				score += group
				group--
			}
		}
		//fmt.Printf("group: %d\n", group)
	}

	return score
}

func read(reader io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	var result []string

	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	return result, scanner.Err()
}
