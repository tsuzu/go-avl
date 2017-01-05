# Go-Avl
- AVL Tree for Golang

# Install
- go get github.com/cs3238-tsuzu/go-Avl

# Usage
- import "github.com/cs3238-tsuzu/go-Avl"
- Creates a new tree with avl.NewTree(comp)
    - The 1st argument of avl.NewTree() is to compare two keys, Comparator.
    - The value Comparator must returns 
        - 1st == 2nd: 0
        - 1st < 2nd: 1
        - 1st > 2nd: -1


# License
- Under the MIT License
- Copyright (c) 2017 Tsuzu