package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

const input = "2112"

func main() {
	first := sumDigitsByStep(input, 1)
	second := sumDigitsByStep(input, len(input)/2)

	fmt.Printf("first: %d, second: %d\n", first, second)

}

func sumDigitsByStep(input string, step int) (sum int) {
	l := utf8.RuneCountInString(input)
	_, w := utf8.DecodeRuneInString(input)

	for i := 0; i < l; i += w {
		nextIx := i + step*w
		if nextIx > l-1 {
			nextIx = nextIx - l
		}

		first, _ := utf8.DecodeRuneInString(input[i:])
		next, _ := utf8.DecodeRuneInString(input[nextIx:])

		if first == next {
			c, _ := strconv.Atoi(string(first))
			sum += c
		}
	}

	return sum
}
