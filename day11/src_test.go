package main

import (
	"reflect"
	"testing"
)

func Test_move(t *testing.T) {
	type args struct {
		hex  Hex
		move string
	}
	tests := []struct {
		name string
		args args
		want Hex
	}{
		{"test", args{hex: Hex{0, 0}, move: "ne"}, Hex{1, 1}},
		{"test", args{hex: Hex{0, 0}, move: "sw"}, Hex{-1, -1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := applyMove(tt.args.hex, tt.args.move); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("applyMove() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_moves(t *testing.T) {
	type args struct {
		hex   Hex
		moves []string
	}
	tests := []struct {
		name string
		args args
		want Hex
	}{
		{"test", args{hex: Hex{0, 0}, moves: []string{"ne", "ne"}}, Hex{2, -2}},
		{"test", args{hex: Hex{0, 0}, moves: []string{"se", "sw", "se", "sw", "sw"}}, Hex{-1, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := applyMoves(tt.args.hex, tt.args.moves); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("applyMoves() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_distance(t *testing.T) {
	type args struct {
		hex1 Hex
		hex2 Hex
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test", args{hex1: Hex{0, 0}, hex2: applyMoves(Hex{0, 0}, []string{"ne", "ne"})}, 2},
		{"test", args{hex1: Hex{0, 0}, hex2: applyMoves(Hex{0, 0}, []string{"se", "sw", "se", "sw", "sw"})}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := distance(tt.args.hex1, tt.args.hex2); got != tt.want {
				t.Errorf("distance() = %v, want %v", got, tt.want)
			}
		})
	}
}
