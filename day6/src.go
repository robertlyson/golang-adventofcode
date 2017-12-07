package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	parameters := strings.Split("14 0 15 12 11 11 3 5 1 6 8 4 9 1 8 4", " ")
	banks := []int{}

	for _, i := range parameters {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		banks = append(banks, j)
	}

	steps, g := countSteps(banks)

	fmt.Printf("banks: %v", banks)
	fmt.Printf("Steps: %d", steps)

	steps2, _ := countSteps(g)

	fmt.Printf("Steps: %d", steps2)
}

func countSteps(banks []int) (int, []int) {
	m := make(map[string]bool)
	m[fmt.Sprintf("%v", banks)] = true
	steps := 1
	var generated []int

	for {
		generated = nextBlock(banks)
		key := fmt.Sprintf("%v", generated)
		_, found := m[key]
		if found {
			break
		}
		m[key] = true
		steps++
	}

	return steps, generated
}

func nextBlock(banks []int) []int {
	index, max := findMaxWithIndex(banks)

	startIndex := index + 1
	if index == len(banks)-1 {
		startIndex = 0
	}

	banks[index] = 0

	for i := startIndex; max > 0; i++ {
		banks[i]++
		max--

		if i == len(banks)-1 {
			i = -1
		}
	}

	return banks
}

func findMaxWithIndex(array []int) (int, int) {
	index := 0
	max := 0

	for i := 0; i < len(array); i++ {
		if max < array[i] {
			index = i
			max = array[i]
		}
	}

	return index, max
}
