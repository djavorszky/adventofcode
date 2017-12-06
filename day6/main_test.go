package main

import (
	"reflect"
	"testing"
)

func Test_iterate(t *testing.T) {
	type args struct {
		banks []int
	}
	tests := []struct {
		name       string
		args       args
		wantResult []int
	}{
		{"First", args{[]int{0, 2, 7, 0}}, []int{2, 4, 1, 2}},
		{"Second", args{[]int{2, 4, 1, 2}}, []int{3, 1, 2, 3}},
		{"Third", args{[]int{3, 1, 2, 3}}, []int{0, 2, 3, 4}},
		{"Fourth", args{[]int{0, 2, 3, 4}}, []int{1, 3, 4, 1}},
		{"Fifth", args{[]int{1, 3, 4, 1}}, []int{2, 4, 1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := iterate(tt.args.banks); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("iterate() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func Test_getIterationCount(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"First", args{[]int{0, 2, 7, 0}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getIterationCount(tt.args.input); got != tt.want {
				t.Errorf("getIterationCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
