package sort

import (
	"algorithm4/base"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	in := base.Inputs[0:]
	BubbleSort(in)

	if !IsSorted(in) {
		t.Errorf("sorted %v", in)
		t.Errorf("   got %v", base.Inputs)
	}
}

func TestInsertSort(t *testing.T) {
	in := base.Inputs[0:]
	InsertSort(in)

	if !IsSorted(in) {
		t.Errorf("sorted %v", in)
		t.Errorf("   got %v", base.Inputs)
	}
}
