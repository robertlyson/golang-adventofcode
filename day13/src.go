package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type FirewallLayer struct {
	Depth           int
	Range           int
	CurrentPosition int
	MoveBack        bool
}

func (firewallLayer FirewallLayer) String() string {
	return fmt.Sprintf("Depth: %d, Range: %d, Position: %d", firewallLayer.Depth, firewallLayer.Range, firewallLayer.CurrentPosition)
}

func (firewallLayer *FirewallLayer) Move() {
	if firewallLayer.MoveBack {
		firewallLayer.CurrentPosition--
	} else {
		firewallLayer.CurrentPosition++
	}

	if firewallLayer.CurrentPosition == firewallLayer.Range {
		firewallLayer.MoveBack = true
	}
	if firewallLayer.CurrentPosition == 1 {
		firewallLayer.MoveBack = false
	}
}

func (firewallLayer *FirewallLayer) Severity() int {
	return firewallLayer.Depth * firewallLayer.Range
}

func main() {
	firewallLayers := parse("./input.txt")
	maxDepth := maxDepth(firewallLayers)
	position, severity := 0, 0

	for currentDepth := 0; currentDepth <= maxDepth; currentDepth++ {
		layer := findLayer(firewallLayers, currentDepth)

		if layer != nil && layer.CurrentPosition == 1 && layer.Depth == position {
			severity += layer.Severity()
		}

		for i := 0; i < len(firewallLayers); i++ {
			firewallLayers[i].Move()
		}

		position++
	}

	fmt.Printf("Severity: %d\n", severity)
}

func findLayer(layers []FirewallLayer, depth int) *FirewallLayer {
	for _, l := range layers {
		if l.Depth == depth {
			return &l
		}
	}

	return nil
}

func maxDepth(layers []FirewallLayer) int {
	max := 0

	for _, l := range layers {
		if max < l.Depth {
			max = l.Depth
		}
	}

	return max
}

func parse(inputFile string) []FirewallLayer {
	file, err := os.Open(inputFile)

	if err != nil {
		log.Fatal(err)
	}

	fileLines, err := read(file)

	result := make([]FirewallLayer, len(fileLines))

	for i, l := range fileLines {
		tmp := strings.Split(l, ": ")

		depth, err := strconv.Atoi(tmp[0])
		if err != nil {
			panic(err.Error())
		}
		rg, err := strconv.Atoi(tmp[1])
		if err != nil {
			panic(err.Error())
		}
		result[i] = FirewallLayer{Depth: depth, Range: rg, CurrentPosition: 1}
	}

	return result
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
