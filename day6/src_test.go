package main

import (
	"reflect"
	"testing"
)

func Test_countSteps(t *testing.T) {
	type args struct {
		banks []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test", args{banks: []int{0, 2, 7, 0}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSteps(tt.args.banks); got != tt.want {
				t.Errorf("countSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findMaxWithIndex(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{"test", args{array: []int{1, 2, 3}}, 2, 3},
		{"test2", args{array: []int{3, 1, 2}}, 0, 3},
		{"test2", args{array: []int{3, 9, 2}}, 1, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := findMaxWithIndex(tt.args.array)
			if got != tt.want {
				t.Errorf("findMaxWithIndex() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("findMaxWithIndex() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_nextBlock(t *testing.T) {
	type args struct {
		banks []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test", args{banks: []int{0, 2, 7, 0}}, []int{2, 4, 1, 2}},
		{"test", args{banks: []int{2, 4, 1, 2}}, []int{3, 1, 2, 3}},
		{"test", args{banks: []int{3, 1, 2, 3}}, []int{0, 2, 3, 4}},
		{"test", args{banks: []int{0, 2, 3, 4}}, []int{1, 3, 4, 1}},
		{"test", args{banks: []int{1, 3, 4, 1}}, []int{2, 4, 1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextBlock(tt.args.banks); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextBlock() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
