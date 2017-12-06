package main

import (
	"fmt"
)

var (
	input = []int{2, 8, 8, 5, 4, 2, 3, 1, 5, 5, 1, 2, 15, 13, 5, 14}
)

func main() {
	fmt.Println(getIterationCount(input))
}

func getIterationCount(input []int) int {
	var (
		count int
		dict  = make(map[string]int)
	)

	for {
		input = iterate(input)
		count++

		if v, ok := dict[fmt.Sprintf("%v", input)]; ok {
			fmt.Printf("Difference: %v\n", count-v)
			return count
		}

		dict[fmt.Sprintf("%v", input)] = count
	}
}

func iterate(banks []int) []int {
	var (
		result = make([]int, len(banks))
		ix     = 0
		max    = banks[ix]
	)

	for i, v := range banks {
		if v > max {
			ix, max = i, v
		}

		result[i] = v
	}

	banks[ix], result[ix] = 0, 0
	ix++

	for i := 0; i < max; i++ {
		if ix > len(banks)-1 {
			ix = 0
		}

		result[ix]++
		ix++
	}

	return result
}
