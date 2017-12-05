package main

import "testing"

func Test_isValid(t *testing.T) {
	type args struct {
		passphrase string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test1", args{passphrase: "aa bb cc dd ee"}, true},
		{"test2", args{passphrase: "aa bb cc dd aa"}, false},
		{"test3", args{passphrase: "aa bb cc dd aaa"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.passphrase); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
