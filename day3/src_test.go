package main

import (
	"reflect"
	"testing"
)

func Test_spiral(t *testing.T) {
	type args struct {
		n float64
	}
	tests := []struct {
		name string
		args args
		want Point
	}{
		{"12", args{n: 12}, Point{2, 1}},
		{"23", args{n: 23}, Point{0, -2}},
		{"1024", args{n: 1024}, Point{-15, 16}},
		{"277678", args{n: 277678}, Point{212, -263}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := spiral(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("spiral() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_distance(t *testing.T) {
	type args struct {
		p1 Point
		p2 Point
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"12", args{p1: Point{0, 0}, p2: Point{2, 1}}, 3},
		{"23", args{p1: Point{0, 0}, p2: Point{0, -2}}, 2},
		{"1024", args{p1: Point{0, 0}, p2: Point{-15, 16}}, 31},
		{"277678", args{p1: Point{0, 0}, p2: Point{212, -263}}, 475},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := distance(tt.args.p1, tt.args.p2); got != tt.want {
				t.Errorf("distance() = %v, want %v", got, tt.want)
			}
		})
	}
}
