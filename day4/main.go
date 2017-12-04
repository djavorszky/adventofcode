package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func main() {
	b, _ := ioutil.ReadFile("input.txt")

	lines := strings.Split(string(b), "\r\n")

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
		runes := []rune(str)
		sort.Sort(byRunes(runes))

		str = string(runes)

		if _, ok := seen[str]; ok {
			return false
		}

		seen[str] = true
	}

	return true
}

type byRunes []rune

func (s byRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s byRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byRunes) Len() int {
	return len(s)
}
