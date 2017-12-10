package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var (
	registers map[string]int
	delim     = "\r\n"
	highest   = 0
)

func main() {
	registers = make(map[string]int)

	contents, _ := ioutil.ReadFile("input.txt")

	lines := strings.Split(string(contents), delim)

	for _, line := range lines {
		task := parse(line)

		if task.cond.eval() {
			task.op.execute()
		}
	}

	var max int
	for _, reg := range registers {
		if reg > max {
			max = reg
		}
	}

	fmt.Printf("Max: %v, highest values: %v", max, highest)

}

type task struct {
	op   operation
	cond condition
}

type operation struct {
	regName     string
	instruction string
	value       int
}

func (o operation) execute() {
	regValue := registers[o.regName]

	if o.instruction == "inc" {
		regValue += o.value
	} else {
		regValue -= o.value
	}

	registers[o.regName] = regValue

	if regValue > highest {
		highest = regValue
	}
}

type condition struct {
	regName string
	cond    string
	value   int
}

func (c condition) eval() bool {
	regValue := registers[c.regName]

	switch c.cond {
	case "<":
		return regValue < c.value
	case "<=":
		return regValue <= c.value
	case ">":
		return regValue > c.value
	case ">=":
		return regValue >= c.value
	case "==":
		return regValue == c.value
	case "!=":
		return regValue != c.value
	}

	return false
}

func parse(line string) task {
	var t task

	opCond := strings.Split(line, "if")
	op := strings.Split(opCond[0], " ")

	val, _ := strconv.Atoi(op[2])

	t.op = operation{
		regName:     op[0],
		instruction: op[1],
		value:       val,
	}

	cond := strings.Split(opCond[1], " ")
	val, _ = strconv.Atoi(cond[3])

	t.cond = condition{
		regName: cond[1],
		cond:    cond[2],
		value:   val,
	}

	return t
}
