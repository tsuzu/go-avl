package avl

import "testing"
import "math/rand"

type TestType struct {
	A, B int
}

func TestInit(t *testing.T) {
	tree := NewTree(func(a, b interface{}) int {
		A := a.(TestType)
		B := b.(TestType)

		if A.A < B.A {
			return 1
		} else if A.A > B.A {
			return -1
		} else {
			if A.B < B.B {
				return 1
			} else if A.B > B.B {
				return -1
			} else {
				return 0
			}
		}
	})

	if tree.root != nil {
		t.Fatal()
	}
}

func TestInsertAndGet(t *testing.T) {
	tree := NewTree(func(a, b interface{}) int {
		A := a.(TestType)
		B := b.(TestType)

		if A.A < B.A {
			return 1
		} else if A.A > B.A {
			return -1
		} else {
			if A.B < B.B {
				return 1
			} else if A.B > B.B {
				return -1
			} else {
				return 0
			}
		}
	})

	if tree.Insert(TestType{100, 1}, 1) == false {
		t.Fatal("Insert failed(1)")
	}
	if tree.Insert(TestType{10, 10}, 2) == false {
		t.Fatal("Insert failed(2)")
	}
	if tree.Insert(TestType{10, 100}, 4) == false {
		t.Fatal("Insert failed(3)")
	}
	if tree.Insert(TestType{10, 100}, 1) == true {
		t.Fatal("Insert failed(4)")
	}

	if n, ok := tree.Get(TestType{100, 1}); !ok {
		t.Fatal("Get failed(1)")
	} else {
		if n.Val.(int) != 1 {
			t.Fatal("Get failed(1.5)")
		}
	}
	if n, ok := tree.Get(TestType{10, 10}); !ok {
		t.Fatal("Get failed(2)")
	} else {
		if n.Val.(int) != 2 {
			t.Fatal("Get failed(2.5)")
		}
	}
	if n, ok := tree.Get(TestType{10, 100}); !ok {
		t.Fatal("Get failed(3)")
	} else {
		if n.Val.(int) != 4 {
			t.Fatal("Get failed(3.5)")
		}
	}
}

func TestInsertAndIndex(t *testing.T) {
	tree := NewTree(func(a, b interface{}) int {
		A := a.(TestType)
		B := b.(TestType)

		if A.A < B.A {
			return 1
		} else if A.A > B.A {
			return -1
		} else {
			if A.B < B.B {
				return 1
			} else if A.B > B.B {
				return -1
			} else {
				return 0
			}
		}
	})

	if tree.Insert(TestType{100, 1}, 1) == false {
		t.Fatal("Insert failed(1)")
	}
	if tree.Insert(TestType{10, 10}, 2) == false {
		t.Fatal("Insert failed(2)")
	}
	if tree.Insert(TestType{10, 100}, 4) == false {
		t.Fatal("Insert failed(3)")
	}

	if n, ok := tree.Index(2); !ok {
		t.Fatal("Index failed(1)")
	} else {
		if n.Val.(int) != 1 {
			t.Fatal("Index failed(1.5)")
		}
	}
	if n, ok := tree.Index(0); !ok {
		t.Fatal("Index failed(2)")
	} else {
		if n.Val.(int) != 2 {
			t.Fatal("Index failed(2.5)")
		}
	}
	if n, ok := tree.Index(1); !ok {
		t.Fatal("Index failed(3)")
	} else {
		if n.Val.(int) != 4 {
			t.Fatal("Index failed(3.5)")
		}
	}
}

func TestInsertAndSet(t *testing.T) {
	tree := NewTree(func(a, b interface{}) int {
		A := a.(TestType)
		B := b.(TestType)

		if A.A < B.A {
			return 1
		} else if A.A > B.A {
			return -1
		} else {
			if A.B < B.B {
				return 1
			} else if A.B > B.B {
				return -1
			} else {
				return 0
			}
		}
	})

	if tree.Insert(TestType{100, 1}, 1) == false {
		t.Fatal("Insert failed(1)")
	}
	if tree.Insert(TestType{10, 10}, 2) == false {
		t.Fatal("Insert failed(2)")
	}
	if tree.Insert(TestType{10, 100}, 4) == false {
		t.Fatal("Insert failed(3)")
	}
	tree.Set(TestType{10, 100}, 8)
	tree.Set(TestType{10, 10}, 4)
	tree.Set(TestType{100, 1}, 2)

	if n, ok := tree.Get(TestType{100, 1}); !ok {
		t.Fatal("Get failed(1)")
	} else {
		if n.Val.(int) != 2 {
			t.Fatal("Get failed(1.5)")
		}
	}
	if n, ok := tree.Get(TestType{10, 10}); !ok {
		t.Fatal("Get failed(2)")
	} else {
		if n.Val.(int) != 4 {
			t.Fatal("Get failed(2.5)")
		}
	}
	if n, ok := tree.Get(TestType{10, 100}); !ok {
		t.Fatal("Get failed(3)")
	} else {
		if n.Val.(int) != 8 {
			t.Fatal("Get failed(3.5)")
		}
	}
}

func TestInsertAndErase(t *testing.T) {
	tree := NewTree(func(a, b interface{}) int {
		A := a.(TestType)
		B := b.(TestType)

		if A.A < B.A {
			return 1
		} else if A.A > B.A {
			return -1
		} else {
			if A.B < B.B {
				return 1
			} else if A.B > B.B {
				return -1
			} else {
				return 0
			}
		}
	})

	if tree.Insert(TestType{100, 1}, 1) == false {
		t.Fatal("Insert failed(1)")
	}
	if tree.Insert(TestType{10, 10}, 2) == false {
		t.Fatal("Insert failed(2)")
	}
	if tree.Insert(TestType{10, 100}, 4) == false {
		t.Fatal("Insert failed(3)")
	}

	if tree.Erase(TestType{10, 100}) == false {
		t.Fatal("Erase failed(1)")
	}
	if tree.Erase(TestType{10, 100}) == true {
		t.Fatal("Erase failed(2)")
	}

	if n, ok := tree.Get(TestType{100, 1}); !ok {
		t.Fatal("Get failed(1)")
	} else {
		if n.Val.(int) != 1 {
			t.Fatal("Get failed(1.5)")
		}
	}
	if n, ok := tree.Get(TestType{10, 10}); !ok {
		t.Fatal("Get failed(2)")
	} else {
		if n.Val.(int) != 2 {
			t.Fatal("Get failed(2.5)")
		}
	}
	if _, ok := tree.Get(TestType{10, 100}); ok {
		t.Fatal("Get failed(3)")
	}
}

func TestInsertAndRank(t *testing.T) {
	tree := NewTree(func(a, b interface{}) int {
		A := a.(TestType)
		B := b.(TestType)

		if A.A < B.A {
			return 1
		} else if A.A > B.A {
			return -1
		} else {
			if A.B < B.B {
				return 1
			} else if A.B > B.B {
				return -1
			} else {
				return 0
			}
		}
	})

	if tree.Insert(TestType{100, 1}, 1) == false {
		t.Fatal("Insert failed(1)")
	}
	if tree.Insert(TestType{10, 10}, 2) == false {
		t.Fatal("Insert failed(2)")
	}
	if tree.Insert(TestType{10, 100}, 4) == false {
		t.Fatal("Insert failed(3)")
	}

	if n := tree.Rank(TestType{100, 1}); n != 2 {
		t.Fatal("Rank failed(1)")
	}
	if n := tree.Rank(TestType{10, 10}); n != 0 {
		t.Fatal("Rank failed(2)")
	}
	if n := tree.Rank(TestType{10, 100}); n != 1 {
		t.Fatal("Rank failed(3)")
	}
	if n := tree.Rank(TestType{100, 100}); n != -1 {
		t.Fatal("Rank failed(4)")
	}
}

func TestRandomNumbers(t *testing.T) {
	const size = 10000

	tree := NewTree(func(a, b interface{}) int {
		A := a.(int)
		B := b.(int)

		if A < B {
			return 1
		} else if A > B {
			return -1
		}

		return 0
	})

	list := rand.Perm(size)
	m := make(map[int]int)
	for i := 0; i < size; i++ {
		m[list[i]] = i
		tree.Insert(list[i], i)
	}

	for i := 0; i < size; i++ {
		if v, ok := tree.Get(list[i]); ok {
			if v.Val.(int) != m[list[i]] {
				t.Fatalf("Get Failed:{%d, %d} should be {%d %d}\n", list[i], v.Val.(int), list[i], m[list[i]])
			}
		} else {
			t.Fatalf("Get Failed: {%d, %d} is not found\n", list[i], m[list[i]])
		}
	}

	for i := 0; i < size; i++ {
		if v, ok := tree.Index(i); ok {
			if v.Val.(int) != m[i] {
				t.Fatalf("Get Failed:{%d, %d} should be {%d %d}\n", i, v.Val.(int), i, m[i])
			}
		} else {
			t.Fatalf("Get Failed: {%d, %d} is not found\n", i, m[i])
		}
	}

	if tree.Size() != size {
		t.Fatal("Size isn't correct.")
	}

	for i := 0; i < size; i++ {
		if !tree.Erase(i) {
			t.Fatal("Erase failed:", i)
		}
	}

	if tree.Size() != 0 {
		t.Fatal("Size isn't correct.")
	}
}

func BenchmarkInsertLinearNumbers(b *testing.B) {
	b.StopTimer()
	b.ResetTimer()
	tree := NewTree(func(a, b interface{}) int {
		A := a.(int)
		B := b.(int)

		if A < B {
			return 1
		} else if A > B {
			return -1
		}

		return 0
	})
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tree.Insert(i, i)
	}
}

func BenchmarkGetLinearNumbers(b *testing.B) {
	b.StopTimer()
	b.ResetTimer()
	tree := NewTree(func(a, b interface{}) int {
		A := a.(int)
		B := b.(int)

		if A < B {
			return 1
		} else if A > B {
			return -1
		}

		return 0
	})
	for i := 0; i < b.N; i++ {
		tree.Insert(i, i)
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		tree.Get(i)
	}
}

func BenchmarkIndexLinearNumbers(b *testing.B) {
	b.StopTimer()
	b.ResetTimer()
	tree := NewTree(func(a, b interface{}) int {
		A := a.(int)
		B := b.(int)

		if A < B {
			return 1
		} else if A > B {
			return -1
		}

		return 0
	})
	for i := 0; i < b.N; i++ {
		tree.Insert(i, i)
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		tree.Index(i)
	}
}

func BenchmarkEraseLinearNumbers(b *testing.B) {
	b.StopTimer()
	b.ResetTimer()
	tree := NewTree(func(a, b interface{}) int {
		A := a.(int)
		B := b.(int)

		if A < B {
			return 1
		} else if A > B {
			return -1
		}

		return 0
	})
	for i := 0; i < b.N; i++ {
		tree.Insert(i, i)
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		tree.Erase(i)
	}
}

func BenchmarkRankLinearNumbers(b *testing.B) {
	b.StopTimer()
	b.ResetTimer()
	tree := NewTree(func(a, b interface{}) int {
		A := a.(int)
		B := b.(int)

		if A < B {
			return 1
		} else if A > B {
			return -1
		}

		return 0
	})
	for i := 0; i < b.N; i++ {
		tree.Insert(i, i)
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		tree.Rank(i)
	}
}
