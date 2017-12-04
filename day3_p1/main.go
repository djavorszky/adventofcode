package main

import "fmt"

func main() {
	fmt.Println(stepCount(368078))
}

// 37  36  35  34  33  32  31
// 38  17  16  15  14  13  30
// 39  18   5   4   3  12  29
// 40  19   6   1   2  11  28
// 41  20   7   8   9  10  27
// 42  21  22  23  24  25  26
// 43  44  45  46  47  48  49

func getDistanceFromMiddle(number int) int {
	if number == 1 {
		return 0
	}

	squareCount, squareSize := 1, 1
	for {
		if number <= squareSize*squareSize {
			break
		}

		squareSize += 2
		squareCount++
	}

	maxElem := squareSize * squareSize

	area := squareSize*4 - 4

	jump, count := maxElem, 0
	for {
		jump = jump - area/8

		if number > jump {
			break
		}

		count++
	}

	if count%2 == 0 {
		return number - jump
	}

	return (squareCount - 1) - (number - jump)
}

func stepCount(number int) int {
	if number == 1 {
		return 0
	}

	squareCount, squareSize := 1, 1
	for {
		if number <= squareSize*squareSize {
			break
		}

		squareSize += 2
		squareCount++
	}

	distFromMiddle := getDistanceFromMiddle(number)

	return squareCount + distFromMiddle - 1
}
