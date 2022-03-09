package search

import "testing"

func TestBisect_Index(t *testing.T) {
	// 找到最左边的值正好等于 x
	a := []int{1, 2, 3, 3, 3, 5, 6}
	x := 3
	// expect := 2
	i := BisectLeft(a, x, 0, len(a))
	if i != len(a) && a[i] == x {
		t.Logf("Find a[%d] = %d", i, a[i])
	} else {
		t.Errorf("Not found.")
	}
}

func TestBisect_LessThan(t *testing.T) {
	// 找到最右边的值小于 x
	a := []int{1, 2, 3, 3, 3, 5, 6}
	x := 3
	// expect := 2
	i := BisectLeft(a, x, 0, len(a))
	if i == 0 {
		t.Errorf("Not found.")
	} else {
		t.Logf("Find a[%d] = %d", i-1, a[i-1])
	}
}

func TestBisect_GreaterEqual(t *testing.T) {
	// 找到最左边的值大于等于 x
	a := []int{1, 2, 3, 3, 3, 5, 6}
	x := 4
	// expect := 2
	i := BisectLeft(a, x, 0, len(a))
	if i == len(a) {
		t.Errorf("Not found.")
	} else {
		t.Logf("Find a[%d] = %d", i, a[i])
	}
}

func TestBisect_LessEqual(t *testing.T) {
	// 找到最右边的值小于等于 x
	a := []int{1, 2, 3, 3, 3, 5, 6}
	x := 3
	// expect := 2
	i := BisectRight(a, x, 0, len(a))
	if i > 0 {
		t.Logf("Find a[%d] = %d", i-1, a[i-1])
	} else {
		t.Errorf("Not found.")
	}
}

func TestBisect_GreaterThan(t *testing.T) {
	// 找到最左边的值大于 x
	a := []int{1, 2, 3, 3, 3, 5, 6}
	x := 2
	// expect := 2
	i := BisectRight(a, x, 0, len(a))
	if i != len(a) {
		t.Logf("Found a[%d] = %d", i, a[i])
	} else {
		t.Errorf("Not found")
	}
}
