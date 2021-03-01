package sort

import "testing"

type item struct {
	id int
}

type items []item

func (r items) Len() int {
	return len(r)
}

func (r items) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r items) Less(i, j int) bool {
	return r[i].id < r[j].id
}

var inputs = items{{74}, {59}, {238}, {1231313}, {-784}, {9845}}

func TestBubbleSort(t *testing.T) {
	in := inputs[0:]
	BubbleSort(in)

	if !IsSorted(in) {
		t.Errorf("sorted %v", in)
		t.Errorf("   got %v", inputs)
	}
}
