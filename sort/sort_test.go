package sort

import "testing"

func TestSort(t *testing.T) {
	basesort := baseSort{}
	firstsort := firstsort{}

	items := make([]item, 10)
	testSorter(&basesort, items)
	testSorter(&firstsort, items)
}

func TestInput(t *testing.T) {
	basesort := baseSort{}
	input(&basesort)
}
