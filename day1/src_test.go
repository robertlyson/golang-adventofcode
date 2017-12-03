package main

import "testing"

func TestCaptcha(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1122", args{input: "1122"}, 3},
		{"1111", args{input: "1111"}, 4},
		{"1234", args{input: "1234"}, 0},
		{"91212129", args{input: "91212129"}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := captcha(tt.args.input); got != tt.want {
				t.Errorf("captcha() = %v, want %v", got, tt.want)
			}
		})
	}
}
