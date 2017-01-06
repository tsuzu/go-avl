# Go-Avl
- AVL Tree for Golang

# Install
- go get github.com/cs3238-tsuzu/go-avl

# Usage
- import "github.com/cs3238-tsuzu/go-avl"
- Creates a new tree with avl.NewTree(comp)
    - The 1st argument of avl.NewTree() is to compare two keys, Comparator.
    - The value Comparator must returns 
        - 1st == 2nd: 0
        - 1st < 2nd: 1
        - 1st > 2nd: -1

# License
- Under the MIT License
- Copyright (c) 2017 Tsuzu

# Benchmark
```
BenchmarkInsertLinearNumbers-12    	 2000000	       877 ns/op
BenchmarkGetLinearNumbers-12       	 3000000	       639 ns/op
BenchmarkIndexLinearNumbers-12     	10000000	       263 ns/op
BenchmarkEraseLinearNumbers-12     	 3000000	       596 ns/op
BenchmarkRankLinearNumbers-12      	 3000000	       637 ns/op
PASS
ok  	github.com/cs3238-tsuzu/go-avl	34.916s
```
