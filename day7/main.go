package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var delim = "\r\n"

func main() {
	input, _ := ioutil.ReadFile("input.txt")

	root := getRoot(string(input))
	getCombinedWeight(root)

	unbalancedNode := findUnbalancedNode(root)
	//fmt.Printf("%#v\n\n", unbalancedNode)

	fmt.Printf("unbalanced node's name and weight: %s, %d\n",
		unbalancedNode.name, unbalancedNode.weight)

	fmt.Println("unbalanced nodes children and their combined weights:")

	correctWeight := findCorrectWeight(unbalancedNode)

	fmt.Println(correctWeight)
}

func findCorrectWeight(n node) int {
	first := getCombinedWeight(n.children[0])

	correctWeight := getCombinedWeight(n.children[1])
	for _, child := range n.children[1:] {
		w := getCombinedWeight(child)
		if w == first {
			correctWeight = w
			break
		}
	}

	for _, child := range n.children {
		w := getCombinedWeight(child)
		if w != correctWeight {
			offset := correctWeight - w
			return child.weight + offset
		}
	}

	return -1
}

// Unbalanced node is the one who has one child which is heavier or lighter than
// the rest.
func findUnbalancedNode(n node) node {
	for _, child := range n.children {
		if !balanced(child) {
			return findUnbalancedNode(child)
		}
	}

	return n
}

type node struct {
	name     string
	weight   int
	parent   *node
	children []node
}

func getRoot(input string) (root node) {
	tree := assembleTree(input)

	for _, node := range tree {
		if node.parent == nil {
			root = *node
			break
		}
	}

	return root
}

func assembleTree(input string) (tree map[string]*node) {
	tree = make(map[string]*node)

	lines := strings.Split(input, delim)

	for _, line := range lines {
		node := nodeFromLine(line)

		tree[node.name] = &node
	}

	for _, node := range tree {
		for i, childNode := range node.children {
			actualNode := tree[childNode.name]

			actualNode.parent = node

			tree[childNode.name] = actualNode
			node.children[i] = *actualNode
		}
	}

	return tree
}

func nodeFromLine(line string) (n node) {
	splitLine := strings.Split(line, "->")

	if len(splitLine) > 1 {
		children := strings.Split(splitLine[1], ",")

		n.children = make([]node, 0)

		for _, child := range children {
			n.children = append(n.children, node{name: strings.Trim(child, " ")})
		}
	}

	nameAndWeight := strings.Split(splitLine[0], " ")
	n.name = nameAndWeight[0]
	weight := nameAndWeight[1]

	w, err := strconv.Atoi(weight[1 : len(weight)-1])
	if err != nil {
		fmt.Printf("Err. %v", err)
	}

	n.weight = w

	return n
}

var weights = make(map[string]int)

func getCombinedWeight(n node) int {
	if w, ok := weights[n.name]; ok {
		return w
	}

	w := n.weight
	for _, child := range n.children {
		w += getCombinedWeight(child)
	}

	weights[n.name] = w

	return w
}

func balanced(n node) bool {
	w := getCombinedWeight(n.children[0])
	for _, child := range n.children {
		if getCombinedWeight(child) != w {
			return false
		}
	}

	return true
}
