package main

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func Test_steps(t *testing.T) {
	type args struct {
		instructions []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{instructions: []int{0, 3, 0, 1, -3}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := steps(tt.args.instructions); got != tt.want {
				t.Errorf("steps() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_read(t *testing.T) {
	type args struct {
		reader io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{"test", args{reader: strings.NewReader("1\r\n-1")}, []int{1, -1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := read(tt.args.reader)
			if (err != nil) != tt.wantErr {
				t.Errorf("read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("read() = %v, want %v", got, tt.want)
			}
		})
	}
}
