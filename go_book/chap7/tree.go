// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 101.

// Package treesort provides insertion sort using an unbalanced binary tree.
package main

import (
	"fmt"
	"math/rand"
)

//!+
type tree struct {
	value       int
	left, right *tree
}

// Sort sorts values in place.
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func (tr *tree) String() string {
	var str string
	var scan func(t *tree)

	scan = func(t *tree) {
		if t.left != nil {
			scan(t.left)
		}
		str = fmt.Sprintf("%s %d", str, t.value)
		if t.right != nil {
			scan(t.right)
		}
	}
	scan(tr)
	return str

}

func main() {

	data := make([]int, 40)

	for i, _ := range data {
		data[i] = rand.Int() % 100
	}
	fmt.Println("unsorted")
	fmt.Println(data)

	var root *tree
	for _, v := range data {
		root = add(root, v)
	}
	fmt.Println("sorted")
	fmt.Println(root.String())

}

//!-
