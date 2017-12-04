package main

import "testing"

func Test_minmax(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{"test", args{input: "1;2;3"}, 1, 3},
		{"test2", args{input: "3;2;1"}, 1, 3},
		{"test3", args{input: "2;3;1"}, 1, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := minmax(tt.args.input)
			if got != tt.want {
				t.Errorf("minmax() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("minmax() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
