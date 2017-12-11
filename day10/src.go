package main

import (
	"fmt"
	"math"
)

func main() {
	array := make([]int, 256)
	for i := 0; i < len(array); i++ {
		array[i] = i
	}
	fmt.Printf("input: %v\n", array)
	hash := hash(array, []int{46, 41, 212, 83, 1, 255, 157, 65, 139, 52, 39, 254, 2, 86, 0, 204})
	fmt.Printf("hash1: %v hash2: %v value: %v\n", hash[0], hash[1], hash[0]*hash[1])
}

func hash(array []int, lenghts []int) []int {
	position := 0
	skip := 0

	for _, l := range lenghts {
		if l > 0 {
			end := ringPosition(len(array), position+l) - 1
			reverse(array, position, end)
		}

		position = ringPosition(len(array), position+l+skip)
		skip++
	}

	return array
}

func ringPosition(lenght int, position int) int {
	if position < lenght {
		return position
	}
	return int(math.Abs(float64(position-lenght))) % lenght
}

func reverse(array []int, start int, end int) {
	if start == end {
		return
	}
	indices := make([]int, 0)
	distance := 0
	if start < end {
		distance = end - start
		for i := 0; i <= distance; i++ {
			indices = append(indices, i+start)
		}
	}
	if start > end {
		distance = len(array) - start + end
		for i, s := 0, start; i <= distance; i, s = i+1, s+1 {
			if s == len(array) {
				s = 0
			}
			indices = append(indices, s)
		}
	}

	for i, j := 0, len(indices)-1; i < j; i, j = i+1, j-1 {
		array[indices[i]], array[indices[j]] = array[indices[j]], array[indices[i]]
	}
}
