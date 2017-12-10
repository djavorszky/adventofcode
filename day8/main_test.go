package main

import (
	"reflect"
	"testing"
)

func Test_parse(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want task
	}{
		{"a", args{"b inc 5 if a > 1"}, task{operation{"b", "inc", 5}, condition{"a", ">", 1}}},
		{"b", args{"a inc 1 if b < 5"}, task{operation{"a", "inc", 1}, condition{"b", "<", 5}}},
		{"c", args{"c dec -10 if a >= 1"}, task{operation{"c", "dec", -10}, condition{"a", ">=", 1}}},
		{"d", args{"c inc -20 if c == 10"}, task{operation{"c", "inc", -20}, condition{"c", "==", 10}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parse(tt.args.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
