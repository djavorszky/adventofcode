package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	b, _ := ioutil.ReadFile("input.txt")

	strs := strings.Split(string(b), "\n")

	instructions := make([]int, len(strs))

	for i, str := range strs {
		num, _ := strconv.Atoi(str)

		instructions[i] = num
	}

	fmt.Println(walkV2(instructions))

}

func walk(instructions []int) int {
	pos := 0

	out, count := len(instructions)-1, 1
	for {
		inst := instructions[pos]
		instructions[pos]++

		next := pos + inst
		if next < 0 || next > out {
			break
		}

		pos = next

		count++
	}

	return count
}

func walkV2(instructions []int) int {
	pos := 0

	out, count := len(instructions)-1, 1
	for {
		inst := instructions[pos]

		next := pos + inst
		if next < 0 || next > out {
			break
		}

		if inst > 2 {
			instructions[pos]--
		} else {
			instructions[pos]++
		}

		pos = next

		count++
	}

	return count
}
