package main

import (
	"reflect"
	"testing"
)

func Test_hash(t *testing.T) {
	type args struct {
		array   []int
		lenghts []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test", args{array: []int{0, 1, 2, 3, 4}, lenghts: []int{3, 4, 1, 5}}, []int{3, 4, 2, 1, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hash(tt.args.array, tt.args.lenghts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("hash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ringPosition(t *testing.T) {
	type args struct {
		lenght   int
		position int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test", args{lenght: 5, position: 1}, 1},
		{"test", args{lenght: 5, position: 6}, 1},
		{"test", args{lenght: 5, position: 5}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ringPosition(tt.args.lenght, tt.args.position); got != tt.want {
				t.Errorf("ringPosition() = %v, want %v", got, tt.want)
			}
		})
	}
}
