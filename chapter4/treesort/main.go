package main

import (
	"fmt"
	"math/rand"
)

type node struct {
	val        int
	leftChild  *node
	rightChild *node
}

func insert(subRoot *node, val int) *node {
	if subRoot == nil {
		newRoot := new(node)
		newRoot.val = val
		return newRoot
	} else if subRoot.val > val {
		subRoot.leftChild = insert(subRoot.leftChild, val)
	} else {
		subRoot.rightChild = insert(subRoot.rightChild, val)
	}
	return subRoot
}

func visit(subRoot *node) []int {
	var result []int
	if subRoot == nil {
		return []int{}
	} else {
		part1 := visit(subRoot.leftChild)
		part2 := visit(subRoot.rightChild)
		result = append(result, part1...)
		result = append(result, subRoot.val)
		result = append(result, part2...)
		return result
	}
}

func treeSort(val []int) []int {
	var root *node
	for _, v := range val {
		root = insert(root, v)
	}
	result := visit(root)
	return result
}

func main() {
	var a [100]int
	for i := range a {
		a[i] = rand.Intn(100)
	}
	fmt.Println(a)
	fmt.Println(treeSort(a[:]))
}
