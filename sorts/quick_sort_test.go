package sorts
import (
	"testing"
	"reflect"
)
func TestQuickSort(t *testing.T) {
	cs := [][][]int{
		{{1}, {1}},
		{{2, 3, 1}, {1, 2, 3}},
		{{2, 5, 4, 9, 3, 8, 7, 1, 6}, {1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{{6, 7, 3, 9, 4, 8, 5, 1, 2}, {1, 2, 3, 4, 5, 6, 7, 8, 9}},
	}

	for _, c := range cs {
		QuickSort(c[0])
		if !reflect.DeepEqual(c[0], c[1]) {
			t.Fatal()
		}
	}
}
