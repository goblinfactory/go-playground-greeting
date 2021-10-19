package learninggo

import (
	"fmt"
	"strings"
)

// IntTree ...
type IntTree struct {
	val         int
	left, right *IntTree
	a1          string
}

// TestBinaryTree1 ...
func TestBinaryTree1() {

	var bt *IntTree
	bt = bt.Insert(10)
	bt.Insert(1)
	bt.Insert(11)
	bt.Insert(15)
	bt.Insert(4)
	bt.Insert(2)
	bt.Insert(50)
	bt.Insert(40)
	bt.Insert(3)
	bt.Dump()
	fmt.Println("")
	fmt.Println(9, bt.Contains(9))
	fmt.Println(4, bt.Contains(40))
	fmt.Println(4, bt.Contains(2))
}

// Insert ...
func (it *IntTree) Insert(val int) *IntTree {
	if it == nil {
		return &IntTree{val: val}
	}
	if val < it.val {
		it.left = it.left.Insert(val)
	} else if val > it.val {
		it.right = it.right.Insert(val)
	}
	return it
}

// Contains ...
func (it *IntTree) Contains(val int) bool {
	fmt.Print("x")
	if it == nil {
		return false
	}
	if val < it.val {
		return it.left.Contains(val)
	} else if val > it.val {
		return it.right.Contains(val)
	}
	return true
}

// Dump ...
func (it *IntTree) Dump() {
	fmt.Println("---")
	it.dump(0, "")
	fmt.Println("---")
}
func (it *IntTree) dump(indent int, you string) {
	if it == nil {
		fmt.Println(you, "nil")
		return
	}
	fmt.Println(strings.Repeat(" ", indent), "-", you, it.val)
	if it.left != nil {
		it.left.dump(indent+2, "L")
	}
	if it.right != nil {
		it.right.dump(indent+2, "R")
	}
}
