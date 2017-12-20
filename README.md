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
goos: darwin
goarch: amd64
pkg: github.com/cs3238-tsuzu/go-avl
BenchmarkInsertLinearNumbers-12    	   10000	   3267119 ns/op
BenchmarkGetLinearNumbers-12       	   10000	   1424781 ns/op
BenchmarkIndexLinearNumbers-12     	   10000	    899655 ns/op
BenchmarkEraseLinearNumbers-12     	   10000	   1868449 ns/op
BenchmarkRankLinearNumbers-12      	   10000	   1511353 ns/op
PASS
ok  	github.com/cs3238-tsuzu/go-avl	221.613s
```
