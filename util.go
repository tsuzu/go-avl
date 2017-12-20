package avl

func maxInt(a, b int) int {
	if a > b {
		return a
	}

	return b
}

var ComparatorInt = func(a, b interface{}) int {
	A := a.(int)
	B := b.(int)

	if A < B {
		return 1
	} else if A > B {
		return -1
	}

	return 0
}

var ComparatorInt64 = func(a, b interface{}) int {
	A := a.(int64)
	B := b.(int64)

	if A < B {
		return 1
	} else if A > B {
		return -1
	}

	return 0
}

var ComparatorFloat32 = func(a, b interface{}) int {
	A := a.(float32)
	B := b.(float32)

	if A < B {
		return 1
	} else if A > B {
		return -1
	}

	return 0
}

var ComparatorFloat64 = func(a, b interface{}) int {
	A := a.(float64)
	B := b.(float64)

	if A < B {
		return 1
	} else if A > B {
		return -1
	}

	return 0
}
