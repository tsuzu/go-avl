package main

import (
	"fmt"

	"github.com/cs3238-tsuzu/go-avl"
)

func main() {
	tree := avl.NewTree(func(a, b interface{}) int {
		A := a.(int)
		B := b.(int)

		if A < B {
			return 1
		} else if A > B {
			return -1
		}

		return 0
	})

	tree.Insert(100, 1)
	if !tree.Insert(100, 10) { // NOT inserted and updated
		fmt.Println("{100, *} is already inserted.")
	}

	tree.Set(100, 10) // Updated

	tree.Set(1, 2)     // Inserted
	tree.Set(3, 4)     // Inserted
	tree.Set(5, 6)     // Inserted
	tree.Set(10, 1000) // Inserted

	fmt.Println("Size:", tree.Size()) //  The size of the tree is returned. (output: "Size: 2")

	if tree.Erase(10) {
		fmt.Println("{10, 1000} has been erased!")
	}

	if !tree.Erase(10) {
		fmt.Println("{10, *} is not inserted.")
	}

	if n, ok := tree.Index(1); ok { // 0-indexed
		fmt.Printf("The 1st node is {%d, %d}.\n", n.Key.(int), n.Val.(int))
	} else {
		fmt.Println("The 1st node is not found.")
	}

	if n := tree.Rank(3); n != -1 {
		fmt.Printf("{3, 4} is the %dst/nd/rd/th node.\n", n) // 0-indexed
	} else {
		fmt.Println("{3, *} is not found.")
	}

	if n, ok := tree.Get(5); ok { // Find with key
		fmt.Println("The value of the node whose key is 5 is", n.Val.(int))
	} else {
		fmt.Println("{10, *} is not found.")
	}
}
