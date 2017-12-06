package main

import (
	"testing"
)

func Test_walk(t *testing.T) {
	type args struct {
		instructions []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Given Example", args{[]int{0, 3, 0, 1, -3}}, 5},
		{"Small Example", args{[]int{0, 1, -1}}, 5},
		{"Negative Example", args{[]int{0, -2}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := walk(tt.args.instructions); got != tt.want {
				t.Errorf("walk() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_walkV2(t *testing.T) {
	type args struct {
		instructions []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Given Example", args{[]int{0, 3, 0, 1, -3}}, 10},
		{"Small Example", args{[]int{0, 1, -1}}, 5},
		{"Negative Example", args{[]int{0, -2}}, 3},
		{"Another Example", args{[]int{0, 3, -1, -1, -1}}, 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := walkV2(tt.args.instructions); got != tt.want {
				t.Errorf("walkV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
