package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

type Hex struct {
	X int
	Y int
}

type Cube struct {
	X int
	Y int
	Z int
}

func main() {
	start, current := Hex{0, 0}, Hex{0, 0}
	content, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err.Error)
	}

	moves := strings.Split(string(content), ",")

	maxDistance := 0
	fmt.Printf("Moves: %v\n", moves)
	for _, move := range moves {
		current = applyMove(current, move)
		currentDistance := distance(start, current)
		if maxDistance < currentDistance {
			maxDistance = currentDistance
		}
	}
	end := applyMoves(start, moves)
	distance := distance(start, end)
	fmt.Printf("Distance: %v\n", distance)
	fmt.Printf("MaxDistance: %v\n", maxDistance)
}

//https://www.redblobgames.com/grids/hexagons/
func distance(a Hex, b Hex) int {
	return int((math.Abs(float64(b.X)) + math.Abs(float64(b.Y))) / 2)
	//return hexDistance(a, b)
}

func axialToCube(hex Hex) Cube {
	var x = hex.X
	var z = hex.Y
	var y = -x - z
	return Cube{x, y, z}
}

func cubeDistance(a Cube, b Cube) int {
	aa := math.Abs(float64(a.X - b.X))
	bb := math.Abs(float64(a.Y - b.Y))
	cc := math.Abs(float64(a.Z - b.Z))

	return int((aa + bb + cc) / 2)
}

func hexDistance(a Hex, b Hex) int {
	ac := axialToCube(a)
	bc := axialToCube(b)
	return cubeDistance(ac, bc)
}

func applyMoves(hex Hex, moves []string) Hex {
	for _, m := range moves {
		hex = applyMove(hex, m)
	}
	return hex
}

func applyMove(hex Hex, move string) Hex {
	x := hex.X
	y := hex.Y

	switch move {
	case "n":
		y--
		break
	case "ne":
		x++
		y--
		break
	case "se":
		x++
		break
	case "s":
		y++
		break
	case "sw":
		x--
		y++
		break
	case "nw":
		x--
		break
	default:
		panic("Unsupported move: " + move + ".")
	}

	return Hex{x, y}
}
