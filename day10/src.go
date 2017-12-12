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

	fmt.Printf("l: %v\n", lenghts2())
	h := []int{}
	lenghts := lenghts2()
	for i := 0; i < 64; i++ {
		h = hash(array, lenghts)
	}

	fmt.Printf("%v\n", h)
	sparseHash := sparseHash(h)

	fmt.Printf("sparseHash: %v\n", sparseHash)

	fmt.Printf("hexadecimal: ")
	for i := 0; i < len(sparseHash); i++ {
		fmt.Printf("%x", sparseHash[i])
	}
}

func sparseHash(array []int) []int {
	sparseHash := make([]int, 16)
	for i := 0; i < 16; i++ {
		for j := i * 16; j < 256; j++ {
			sparseHash[i] ^= array[j]
		}
	}

	return sparseHash
}

func lenghts() []int {
	lenghts := []int{46, 41, 212, 83, 1, 255, 157, 65, 139, 52, 39, 254, 2, 86, 0, 204}
	return lenghts
}

func lenghts2() []int {
	input := "46,41,212,83,1,255,157,65,139,52,39,254,2,86,0,204"
	lenghts := make([]int, 0)
	for i := 0; i < len(input); i++ {
		lenghts = append(lenghts, int(input[i]))
	}

	lenghts = append(lenghts, 17, 31, 73, 47, 23)
	return lenghts
}

var position = 0
var skip = 0

func hash(input []int, lenghts []int) []int {
	array := make([]int, len(input))
	copy(array, input)

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
