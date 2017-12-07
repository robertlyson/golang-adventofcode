package main

import (
	"reflect"
	"testing"
)

func Test_parseLine(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want Line
	}{
		{"test", args{line: "fwft (72) -> ktlj, cntj, xhth"}, Line{"fwft", []string{"ktlj", "cntj", "xhth"}}},
		{"test", args{line: "meeiw (95)"}, Line{"meeiw", []string{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseLine(tt.args.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findRoot(t *testing.T) {
	type args struct {
		lines []Line
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test", args{lines: []Line{Line{"a", []string{"b,c,d"}}, Line{"b", []string{"c", "d"}}}}, "a"},
		{"test", args{lines: []Line{Line{"a", []string{"c,d"}}, Line{"b", []string{"c", "a"}}}}, "b"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRoot(tt.args.lines); got != tt.want {
				t.Errorf("findRoot() = %v, want %v", got, tt.want)
			}
		})
	}
}
