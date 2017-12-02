package main

import (
	"fmt"
	"strconv"
	"strings"
)

var inputs = []string{
	"123	456",
	"789	123	8182	821	4	1822",
}

func main() {
	var checksum int
	for _, row := range inputs {
		entries := strings.Split(row, "\t")
		ints := stringsToInts(entries)
		checksum += rowChecksum(ints)
	}

	fmt.Println(checksum)
}

func stringsToInts(strs []string) []int {
	var res []int
	for _, str := range strs {
		i, _ := strconv.Atoi(str)

		res = append(res, i)
	}

	return res
}

func rowChecksum(nums []int) int {
	var min, max = nums[0], nums[0]

	for _, num := range nums {
		if num < min {
			min = num
		} else if num > max {
			max = num
		}
	}

	return max - min
}
