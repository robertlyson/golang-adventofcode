package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	instructions, err := read(file)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("instructions loaded: ", instructions)

	steps := steps(instructions)

	fmt.Printf("steps: %d", steps)
}

func steps(instructions []int) int {
	steps, position := 0, 0

	for position < len(instructions) {
		instruction := &instructions[position]
		position += *instruction
		*instruction++
		steps++
	}

	return steps
}

func read(reader io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	var result []int

	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return result, err
		}
		result = append(result, x)
	}

	return result, scanner.Err()
}
