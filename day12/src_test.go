package main

import (
	"testing"
)

func Test_connections(t *testing.T) {
	type args struct {
		programsDefinition []string
		programID          string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test", args{programsDefinition: []string{"1 <-> ", "2 <-> 1,2"}, programID: "1"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := connections(tt.args.programsDefinition, tt.args.programID); got != tt.want {
				t.Errorf("connections() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_do(t *testing.T) {
	type args struct {
		inputFile string
	}
	tests := []struct {
		name string
		args args
	}{
		{"test", args{inputFile: "./input.txt"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			do(tt.args.inputFile)
		})
	}
}
