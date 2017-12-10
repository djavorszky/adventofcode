package main

import (
	"testing"
)

func Test_removeGarbage(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"empty group", args{"{}"}, "{}"},
		{"group with gb", args{"{<a>}"}, "{}"},
		{"group with more gb", args{"{<{},{},{{}}>}"}, "{}"},
		{"group with gb + escape", args{"{<a!>},{<a!>},{<a!>},{<ab>}"}, "{}"},
		{"group with multiple gb", args{"{<a>,<a>,<a>,<a>}"}, "{,,,}"},
		{"multi-escapes", args{"{{<!!>},{<!!>},{<!!>},{<!!>}}"}, "{{},{},{},{}}"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeGarbage(tt.args.input); got != tt.want {
				t.Errorf("removeGarbage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countScore(t *testing.T) {
	type args struct {
		cleanedInput string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"empty", args{"{}"}, 1},
		{"3deep", args{"{{{}}}"}, 6},
		{"nested", args{"{{{},{},{{}}}}"}, 16},
		{"stuff", args{"{{},{},{},{}}"}, 9},
		{"commas", args{"{,,,,}"}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countScore(tt.args.cleanedInput); got != tt.want {
				t.Errorf("countScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
