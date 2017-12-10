package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"unicode/utf8"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")

	cleanedInput, removed := removeGarbage(string(input))

	fmt.Printf("Removed: %d, score: %d", removed, countScore(cleanedInput))

}

func countScore(cleanedInput string) int {
	var (
		w      int
		count  int
		result int
		cur    rune
	)

	l := utf8.RuneCountInString(cleanedInput)

	for i := 0; i < l; i += w {
		cur, w = utf8.DecodeRuneInString(cleanedInput[i:])

		if cur == '{' {
			count++
			result += count
		} else if cur == '}' {
			count--
		}

	}

	return result
}

func removeGarbage(input string) (string, int) {
	var (
		result    bytes.Buffer
		w         int
		prev, cur rune
		removed   int
	)

	garbage := false
	l := utf8.RuneCountInString(input)

	for i := 0; i < l; i += w {
		cur, w = utf8.DecodeRuneInString(input[i:])

		if !garbage {
			if prev != '!' && cur == '<' {
				garbage = true
				continue
			}

			if prev != '!' && cur == '>' {
				garbage = false
				continue
			}

			result.WriteRune(cur)

			prev = cur
		} else {
			if cur == '!' {
				i += w
				continue
			}

			if cur == '>' {
				garbage = false
				continue
			}

			removed++
		}
	}

	return result.String(), removed
}
