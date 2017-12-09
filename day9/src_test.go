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
			if got := parse(tt.args.input); got != tt.want {
				t.Errorf("parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
