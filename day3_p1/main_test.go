package main

import (
	"testing"
)

func Test_stepCount(t *testing.T) {
	type args struct {
		number int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{1}, 0},
		{"12", args{12}, 3},
		{"23", args{23}, 2},
		{"1024", args{1024}, 31},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stepCount(tt.args.number); got != tt.want {
				t.Errorf("stepCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getDistanceFromMiddle(t *testing.T) {
	type args struct {
		number int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{1}, 0},
		{"2", args{2}, 0},
		{"3", args{3}, 1},
		{"4", args{4}, 0},
		{"5", args{5}, 1},
		{"6", args{6}, 0},
		{"7", args{7}, 1},
		{"8", args{8}, 0},
		{"9", args{9}, 1},
		{"10", args{10}, 1},
		{"11", args{11}, 0},
		{"12", args{12}, 1},
		{"13", args{13}, 2},
		{"14", args{14}, 1},
		{"15", args{15}, 0},
		{"16", args{16}, 1},
		{"17", args{17}, 2},
		{"18", args{18}, 1},
		{"19", args{19}, 0},
		{"20", args{20}, 1},
		{"21", args{21}, 2},
		{"22", args{22}, 1},
		{"23", args{23}, 0},
		{"24", args{24}, 1},
		{"25", args{25}, 2},
		{"26", args{26}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDistanceFromMiddle(tt.args.number); got != tt.want {
				t.Errorf("getDistanceFromMiddle() = %v, want %v", got, tt.want)
			}
		})
	}
}
