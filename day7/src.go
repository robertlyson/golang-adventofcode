package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Node struct {
	Name          string
	Parent        *Node
	ChildrenNames map[string]bool
	Children      []*Node
	Weight        int
}

func (n *Node) CalcWeight() int {
	if n.Children == nil {
		return n.Weight
	}

	weight := n.Weight

	for i := range n.Children {
		weight += n.Children[i].CalcWeight()
	}

	return weight
}

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	fileLines, err := read(file)
	nodes := make([]Node, len(fileLines))
	for i := 0; i < len(nodes); i++ {
		nodes[i] = parseLine(fileLines[i])
	}

	root := findRoot(nodes)
	fmt.Printf("Root: %v\n", root.Name)

	for i := range root.Children {
		weight := root.Children[i].CalcWeight()
		fmt.Printf("node: %v weight: %d weight: %d\n", root.Children[i].Name, root.Children[i].Weight, weight)
	}
}

func findRoot(lines []Node) Node {
	for i := 0; i < len(lines); i++ {
		current := &lines[i]

		for j := 0; j < len(lines); j++ {
			_, found := lines[j].ChildrenNames[current.Name]
			if found {
				current.Parent = &lines[j]
				lines[j].Children = append(lines[j].Children, current)
				break
			}
		}
	}

	root := Node{}

	for i := 0; i < len(lines); i++ {
		if lines[i].Parent == nil {
			root = lines[i]
			break
		}
	}

	return root
}

func parseLine(line string) Node {
	parsed := strings.Split(line, " -> ")

	r, _ := regexp.Compile("(.+) \\((\\d+)\\)")

	groups := r.FindStringSubmatch(parsed[0])

	children := make(map[string]bool)
	if len(parsed) > 1 {
		arr := strings.Split(parsed[1], ", ")
		for s := range arr {
			children[arr[s]] = true
		}
	}
	weight, _ := strconv.Atoi(groups[2])
	return Node{groups[1], nil, children, []*Node{}, weight}
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
