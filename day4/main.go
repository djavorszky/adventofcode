package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	b, _ := ioutil.ReadFile("input.txt")

	lines := strings.Split(string(b), "\r\n")

	fmt.Println(lines[len(lines)-1])

	valid := 0
	for _, line := range lines {
		if isValid(line) {
			valid++
		}
	}

	fmt.Println(valid)
}

func isValid(pass string) bool {
	seen := make(map[string]bool)

	for _, str := range strings.Split(pass, " ") {
		if _, ok := seen[str]; ok {
			return false
		}

		seen[str] = true
	}

	return true
}
