package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

type Line struct {
	Name     string
	Children []string
}

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	fileLines, err := read(file)
	lines := make([]Line, len(fileLines))
	for i := 0; i < len(lines); i++ {
		lines[i] = parseLine(fileLines[i])
	}

	fmt.Printf("Root: %v", findRoot(lines))
}

func findRoot(lines []Line) string {
	names := make([]string, len(lines))
	children := make(map[string]bool)

	for i := 0; i < len(lines); i++ {
		names[i] = lines[i].Name

		for j := 0; j < len(lines[i].Children); j++ {
			key := lines[i].Children[j]
			_, found := children[key]
			if found {
				continue
			}
			children[key] = true
		}
	}

	for i := 0; i < len(lines); i++ {
		_, found := children[lines[i].Name]
		if found == false {
			return lines[i].Name
		}
	}

	return ""
}

func parseLine(line string) Line {
	parsed := strings.Split(line, " -> ")

	r, _ := regexp.Compile("(.+) \\((\\d+)\\)")

	groups := r.FindStringSubmatch(parsed[0])

	children := []string{}
	if len(parsed) > 1 {
		children = strings.Split(parsed[1], ", ")
	}

	return Line{groups[1], children}
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
