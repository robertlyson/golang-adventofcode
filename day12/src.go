package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	part2("./input.txt")
}

func part2(inputFile string) {
	file, err := os.Open(inputFile)

	if err != nil {
		log.Fatal(err)
	}

	fileLines, err := read(file)

	programs := parseAll(fileLines)

	groups := make(map[string]int, 0)

	for _, p := range programs {
		if len(groups) == 0 {
			groups[p.ID]++
			continue
		}

		notInGroup := true
		for key, _ := range groups {
			if p.hasConnectionTo(key, make([]*Program, 0)) {
				groups[key]++
				notInGroup = false
				break
			}
		}
		if notInGroup {
			groups[p.ID]++
		}
	}

	fmt.Printf("Groups: %d", len(groups))
}

func part1(inputFile string) {
	file, err := os.Open(inputFile)

	if err != nil {
		log.Fatal(err)
	}

	fileLines, err := read(file)

	connections := connections(fileLines, "0")

	fmt.Printf("No. of connections to 0: %d\n", connections)
}

type Program struct {
	ID              string
	ConnectionNames []string
	Connections     []*Program
	Parent          *Program
}

func (p Program) hasParent(programID string) bool {
	if p.Parent == nil {
		return false
	}

	if p.Parent.ID == programID {
		return true
	}

	return p.Parent.hasParent(programID)
}

func (p Program) hasConnectionTo(programID string, checked []*Program) bool {
	for _, checked := range checked {
		if checked.ID == p.ID {
			return false
		}
	}
	if p.ID == programID {
		return true
	}
	checked = append(checked, &p)
	for _, c := range p.Connections {
		if c == nil {
			continue
		}
		if c.hasConnectionTo(programID, checked) {
			return true
		}
	}

	return false
}

func connections(programsDefinition []string, programID string) int {
	connections := 0
	programs := parseAll(programsDefinition)

	for _, p := range programs {
		checked := make([]*Program, 0)
		if p.hasConnectionTo(programID, checked) {
			connections++
		}
	}

	return connections
}

func parseAll(programsDefinition []string) map[string]*Program {
	programs := make(map[string]*Program, len(programsDefinition))

	for i := 0; i < len(programsDefinition); i++ {
		program := parse(programsDefinition[i])
		programs[program.ID] = &program
	}

	for _, p := range programs {
		p.Connections = make([]*Program, len(p.ConnectionNames))
		for i := 0; i < len(p.ConnectionNames); i++ {
			connection := p.ConnectionNames[i]
			if connection == p.ID {
				continue
			}
			program, found := programs[connection]
			if found {
				p.Connections[i] = program
			}
		}
	}

	return programs
}

func parse(input string) Program {
	tmp := strings.Split(input, " <-> ")
	id := tmp[0]
	var names []string
	if tmp[1] == "" {
		names = []string{}
	} else {
		tmp := strings.Split(tmp[1], ",")
		for _, r := range tmp {
			names = append(names, strings.TrimSpace(r))
		}
	}

	return Program{ID: id, ConnectionNames: names, Connections: []*Program{}}
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
