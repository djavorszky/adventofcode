package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

var inputs = []string{
	"3093	749	3469	142	2049	3537	1596	3035	2424	3982	3290	125	249	131	118	3138",
	"141	677	2705	2404	2887	2860	1123	2714	117	1157	2607	1800	153	130	1794	3272",
}

func main() {
	var checksum int
	for _, row := range inputs {
		entries := strings.Split(row, "\t")
		ints := stringsToInts(entries)
		checksum += rowChecksumV2(ints)
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

func rowChecksumV1(nums []int) int {
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

func rowChecksumV2(nums []int) int {
	var divRes float64
	for i, frst := range nums {
		for _, scnd := range nums[i+1:] {
			if frst > scnd {
				divRes = float64(frst) / float64(scnd)
			} else {
				divRes = float64(scnd) / float64(frst)
			}

			if divRes == math.Trunc(divRes) {
				return int(divRes)
			}
		}
	}

	return 0
}
