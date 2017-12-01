package main

import (
	"testing"
)

func Test_sumDigits(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1122", args{"1122"}, 3},
		{"1111", args{"1111"}, 4},
		{"1234", args{"1234"}, 0},
		{"91212129", args{"91212129"}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumDigits(tt.args.input); got != tt.want {
				t.Errorf("sumDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sumDigitsHalfway(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1212", args{"1212"}, 6},
		{"1221", args{"1221"}, 0},
		{"123425", args{"123425"}, 4},
		{"123123", args{"123123"}, 12},
		{"12131415", args{"12131415"}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumDigitsHalfway(tt.args.input); got != tt.want {
				t.Errorf("sumDigitsHalfway() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sumDigitsByStep(t *testing.T) {
	type args struct {
		input string
		step  int
	}
	tests := []struct {
		name    string
		args    args
		wantSum int
	}{
		{"1122", args{"1122", 1}, 3},
		{"1111", args{"1111", 1}, 4},
		{"1234", args{"1234", 1}, 0},
		{"91212129", args{"91212129", 1}, 9},

		{"1212", args{"1212", 2}, 6},
		{"1221", args{"1221", 2}, 0},
		{"123425", args{"123425", 3}, 4},
		{"123123", args{"123123", 3}, 12},
		{"12131415", args{"12131415", 4}, 4}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSum := sumDigitsByStep(tt.args.input, tt.args.step); gotSum != tt.wantSum {
				t.Errorf("sumDigitsByStep() = %v, want %v", gotSum, tt.wantSum)
			}
		})
	}
}
