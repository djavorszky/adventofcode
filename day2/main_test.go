package main

import (
	"reflect"
	"testing"
)

func Test_rowChecksum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test 1", args{[]int{5, 1, 9, 5}}, 8},
		{"Test 2", args{[]int{7, 5, 3}}, 4},
		{"Test 3", args{[]int{2, 4, 6, 8}}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rowChecksumV1(tt.args.nums); got != tt.want {
				t.Errorf("rowChecksum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stringsToInts(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Test 1", args{[]string{"5", "1", "9", "5"}}, []int{5, 1, 9, 5}},
		{"Test 2", args{[]string{"7", "5", "3"}}, []int{7, 5, 3}},
		{"Test 3", args{[]string{"2", "4", "6", "8"}}, []int{2, 4, 6, 8}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stringsToInts(tt.args.strs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("stringsToInts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rowChecksumV2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test 1", args{[]int{5, 9, 2, 8}}, 4},
		{"Test 2", args{[]int{9, 4, 7, 3}}, 3},
		{"Test 3", args{[]int{3, 8, 6, 5}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rowChecksumV2(tt.args.nums); got != tt.want {
				t.Errorf("rowChecksumV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
