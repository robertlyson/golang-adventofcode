package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"

	"github.com/knetic/govaluate"
)

//https://github.com/Knetic/govaluate
type Instruction struct {
	Register          string
	Action            string
	Value             int
	Expression        string
	ConditionRegister string
	Condition         string
	ConditionValue    int
}

type Register struct {
	Name  string
	Value int
}

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	fileLines, err := read(file)
	instructions := make([]Instruction, len(fileLines))

	for i := 0; i < len(instructions); i++ {
		instructions[i] = parse(fileLines[i])
	}

	registers := make([]Register, 0)
	for i := range instructions {
		if findRegister(instructions[i].Register, registers) == nil {
			registers = append(registers, Register{Name: instructions[i].Register, Value: 0})
		}
	}

	maxEver := 0

	for i := range instructions {
		instruction := &instructions[i]

		expr, _ := govaluate.NewEvaluableExpression(instruction.Expression)
		conditionRegister := findRegister(instruction.ConditionRegister, registers)

		parameters := make(map[string]interface{}, 1)
		parameters[conditionRegister.Name] = conditionRegister.Value
		result, err := expr.Evaluate(parameters)
		reg := findRegister(instruction.Register, registers)

		if err == nil {
			if result == true {
				if instruction.Action == "inc" {
					reg.Value += instruction.Value
				}
				if instruction.Action == "dec" {
					reg.Value -= instruction.Value
				}
				if reg.Value > maxEver {
					maxEver = reg.Value
				}
			}
		}
	}

	fmt.Printf("%v\n", registers)

	max := 0
	for i := range registers {
		if registers[i].Value > max {
			max = registers[i].Value
		}
	}

	fmt.Printf("Max: %d\n", max)
	fmt.Printf("MaxEver: %d\n", maxEver)
}

func findRegister(name string, registers []Register) *Register {
	for i := range registers {
		if registers[i].Name == name {
			return &registers[i]
		}
	}

	return nil
}

func parse(str string) Instruction {
	r, _ := regexp.Compile("(.+) (.+) (-?\\d+) if ((.+) (.+) (-?.+))")

	groups := r.FindStringSubmatch(str)
	value, _ := strconv.Atoi(groups[3])
	conditionValue, _ := strconv.Atoi(groups[7])
	return Instruction{Register: groups[1], Action: groups[2], Value: value,
		Expression: groups[4], ConditionRegister: groups[5], Condition: groups[6], ConditionValue: conditionValue}
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
