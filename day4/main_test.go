package main

import "testing"

func Test_isValid(t *testing.T) {
	type args struct {
		pass string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"TestOne", args{"aa bb cc dd ee"}, true},
		{"TestTwo", args{"aa bb cc dd aa"}, false},
		{"TestThree", args{"aa bb cc dd aaa"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.pass); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
