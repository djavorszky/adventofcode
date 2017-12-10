package main

import (
	"reflect"
	"testing"
)

var input = `pbga (66)
xhth (57)
ebii (61)
havc (66)
ktlj (57)
fwft (72) -> ktlj, cntj, xhth
qoyq (66)
padx (45) -> pbga, havc, qoyq
tknk (41) -> ugml, padx, fwft
jptl (61)
ugml (68) -> gyxo, ebii, jptl
gyxo (61)
cntj (57)`

func Test_nodeFromLine(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name  string
		args  args
		wantN node
	}{
		{"simpleNode", args{"ktlj (57)"}, node{name: "ktlj"}},
		{"nodeWithChildren", args{"fwft (72) -> ktlj, cntj, xhth"},
			node{name: "fwft",
				children: []node{
					node{name: "ktlj"},
					node{name: "cntj"},
					node{name: "xhth"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotN := nodeFromLine(tt.args.line); !reflect.DeepEqual(gotN, tt.wantN) {
				t.Errorf("nodeFromLine() = %#v, want %#v", gotN, tt.wantN)
			}
		})
	}
}
