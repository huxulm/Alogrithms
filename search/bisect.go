package search

// 返回在数组 a 中插入 x 的索引，假设 a 已排序
//
// 返回值 i 是使得 a[:i] 中的所有 e 都具有 e < x, 并且所有 e 在 a[i:] 有 e >= x。
//
// 所以:
//
// 1. 如果 x 在数组 a 中, 返回的位置是: 在最左侧 x 前插入的位置
//
// 2. 如果 x 不在数组 a 中, 返回值是 0 或 len(a)
//
// 综合1,2 返回值表示 a 中小于 x 的元素个数
//
// 可选参数 lo（默认 0）和 hi（默认 len(a)）
func BisectLeft(a []int, x, lo, hi int) int {
	for lo < hi {
		mid := int(uint(lo+hi) >> 1)
		if a[mid] < x {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}

// 返回在列表 a 中插入 x 的索引，假设 a 已排序。
//
// 返回值 i 使得 a[:i] 中的所有 e 都具有 e <= x, 并且所有 e 在 a[i:] 有 e > x。
//
// 所以如果 x 已经在 a 中， 返回的位置是: 在最右边的 x 之后插入的位置
func BisectRight(a []int, x, lo, hi int) int {
	for lo < hi {
		mid := int(uint(lo+hi) >> 1)
		if x < a[mid] {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}
