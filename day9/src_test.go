package main

import "testing"

func Test_parse(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{input: "{}"}, 1},
		{"test2", args{input: "{{{}}}"}, 6},
		{"test3", args{input: "{{},{}}"}, 5},
		{"test4", args{input: "{{{},{},{{}}}}"}, 16},
		{"test5", args{input: "{<a>,<a>,<a>,<a>}"}, 1},
		{"test6", args{input: "{{<ab>},{<ab>},{<ab>},{<ab>}}"}, 9},
		{"test7", args{input: "{{<!!>},{<!!>},{<!!>},{<!!>}}"}, 9},
		{"test8", args{input: "{{<a!>},{<a!>},{<a!>},{<ab>}}"}, 3},
		{"test9", args{input: "{<}{>}"}, 1},
		{"test10", args{input: "{{!}},{}}"}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := parse(tt.args.input); got != tt.want {
				t.Errorf("parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parse_garbage(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{input: "<>"}, 0},
		{"test2", args{input: "<random characters>"}, 17},
		{"test3", args{input: "<<<<>"}, 3},
		{"test4", args{input: "<{!>}>"}, 2},
		{"test5", args{input: "<!!>"}, 0},
		{"test6", args{input: "<!!!>>"}, 0},
		{"test7", args{input: "<{o\"i!a,<{i<a>"}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, got := parse(tt.args.input); got != tt.want {
				t.Errorf("parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
