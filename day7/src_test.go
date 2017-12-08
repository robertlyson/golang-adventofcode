package main

import "testing"

func TestNode_CalcWeight(t *testing.T) {
	node1 := Node{Weight: 1}
	node1_1 := Node{Weight: 1}
	node1.Children = append(node1.Children, &node1_1)
	node2 := Node{Weight: 2}
	root := Node{Weight: 1}
	root.Children = append(root.Children, &node1)
	root.Children = append(root.Children, &node2)

	tests := []struct {
		name string
		n    *Node
		want int
	}{
		{"test", &root, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.CalcWeight(); got != tt.want {
				t.Errorf("Node.CalcWeight() = %v, want %v", got, tt.want)
			}
		})
	}
}
