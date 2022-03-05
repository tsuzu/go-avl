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
goarch: arm64
pkg: github.com/tsuzu/go-avl
BenchmarkInsertLinearNumbers-10          4092988               319.3 ns/op
BenchmarkGetLinearNumbers-10             9983720               163.5 ns/op
BenchmarkIndexLinearNumbers-10          13540413               130.3 ns/op
BenchmarkEraseLinearNumbers-10           7530286               203.8 ns/op
BenchmarkRankLinearNumbers-10            8986261               173.8 ns/op
PASS
ok      github.com/tsuzu/go-avl 23.888s
```
