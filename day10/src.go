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
	//fmt.Printf("input: %v\n", array)
	hash := hash(array, []int{46, 41, 212, 83, 1, 255, 157, 65, 139, 52, 39, 254, 2, 86, 0, 204})
	fmt.Printf("hash1: %d hash2: %d value: %d\n", hash[0], hash[1], hash[0]*hash[1])
}

func hash(array []int, lenghts []int) []int {
	position := 0
	skip := 0

	for _, l := range lenghts {
		end := ringPosition(len(array), position+l) - 1
		position = ringPosition(len(array), position)
		reverse(array, position, end)

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

func reverse(s []int, start int, end int) {
	if start < end {
		for i, j := start, end; i < j; i, j = i+1, j-1 {
			s[i], s[j] = s[j], s[i]
		}
	}
	if start > end {
		steps := int((len(s)-start+end)/2) + 1
		for i, j, k := start, end, 0; k < steps; i, j, k = i+1, j-1, k+1 {
			if i == len(s) {
				i = 0
			}
			if j < 0 {
				j = len(s) - 1
			}
			s[i], s[j] = s[j], s[i]
		}
	}
}
