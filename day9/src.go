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

	score, garbage := parse(line)

	fmt.Printf("Score %d Garbage %d\n", score, garbage)
}

func parse(input string) (int, int) {
	score := 0
	group := 0
	garbage := false
	cancel := false
	garbageCount := 0

	for _, char := range input {
		if cancel {
			cancel = false
			continue
		}
		if char == '!' && cancel == false {
			cancel = true
			continue
		}
		if garbage && char != '>' {
			garbageCount++
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
	}

	return score, garbageCount
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
