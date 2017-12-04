package main

import (
	"fmt"
	"strconv"
	"strings"
)

const num = 289326

const (
	right = iota
	up
	left
	down
)

var (
	spiral    = make(map[string]int)
	direction = right
	location  = "0,0"
)

func main() {
	spiral[location] = 1

	for {
		move()
		if res := calculate(); res > num {
			fmt.Println(res)
			break
		}
	}
}

// 147  142  133  122   59
// 304    5    4    2   57
// 330   10    1    1   54
// 351   11   23   25   26
// 362  747  806--->   ...

func calculate() int {
	loc := strings.Split(location, ",")
	x, _ := strconv.Atoi(loc[0])
	y, _ := strconv.Atoi(loc[1])

	dirs := []string{
		fmt.Sprintf("%d,%d", x-1, y-1),
		fmt.Sprintf("%d,%d", x-1, y),
		fmt.Sprintf("%d,%d", x-1, y+1),
		fmt.Sprintf("%d,%d", x, y-1),
		fmt.Sprintf("%d,%d", x, y+1),
		fmt.Sprintf("%d,%d", x+1, y-1),
		fmt.Sprintf("%d,%d", x+1, y),
		fmt.Sprintf("%d,%d", x+1, y+1),
	}

	var res int
	for _, dir := range dirs {
		if v, ok := spiral[dir]; ok {
			res += v
		}
	}

	spiral[location] = res

	return res
}

func move() {
	loc := strings.Split(location, ",")
	x, _ := strconv.Atoi(loc[0])
	y, _ := strconv.Atoi(loc[1])

	switch direction {
	case left:
		location = fmt.Sprintf("%d,%d", x-1, y)
	case down:
		location = fmt.Sprintf("%d,%d", x, y-1)
	case right:
		location = fmt.Sprintf("%d,%d", x+1, y)
	case up:
		location = fmt.Sprintf("%d,%d", x, y+1)
	}

	if shouldTurn() {
		turn()
	}
}

func shouldTurn() bool {
	loc := strings.Split(location, ",")
	x, _ := strconv.Atoi(loc[0])
	y, _ := strconv.Atoi(loc[1])

	var checkLock string
	switch direction {
	case up:
		checkLock = fmt.Sprintf("%d,%d", x-1, y)
	case left:
		checkLock = fmt.Sprintf("%d,%d", x, y-1)
	case down:
		checkLock = fmt.Sprintf("%d,%d", x+1, y)
	case right:
		checkLock = fmt.Sprintf("%d,%d", x, y+1)
	}

	_, ok := spiral[checkLock]

	return !ok
}

func turn() {
	switch direction {
	case right:
		direction = up
	case up:
		direction = left
	case left:
		direction = down
	case down:
		direction = right
	}
}
