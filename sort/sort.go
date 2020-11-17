package sort

import (
	"fmt"
)

type item struct {
	item string
}

type sorter interface {
	sort([]item)
	less(item, item) bool
	exch([]item, int, int)
	show([]item)
	isSorted(sorter, []item) bool
}

type baseSort struct{}

func (s *baseSort) show(items []item) {
	fmt.Println("show item:", items)
}

func (s *baseSort) isSorted(sort sorter, items []item) bool {
	for i := 1; i < len(items); i++ {
		if sort.less(items[i], items[i-1]) {
			return false
		}
	}
	return true
}

func (s *baseSort) sort(items []item) {
	fmt.Println("no imp sort")
}

func (s *baseSort) exch(items []item, i int, j int) {
	fmt.Println("no imp exch")
}

func (s *baseSort) less(a1 item, a2 item) bool {
	return a1.item < a2.item
}

type firstsort struct {
	baseSort
}

func (f *firstsort) show(items []item) {
	fmt.Println("show first item:", items)
}

func (f *firstsort) less(a1 item, a2 item) bool {
	fmt.Println("first less")
	return true
}

func testSorter(base sorter, items []item) {
	base.show(items)
	base.isSorted(base, items)
}

func input(base sorter) {
	fmt.Println("res=", S)
	testSorter(base, S)
}

type sorter1 interface {
	sort([]item)
	less(item, item) bool
}

func exch(items []item, i int, j int) {
	tmp := items[j]
	items[j] = items[i]
	items[i] = tmp
}

func show(items []item) {
	fmt.Println("show item:", items)
}

func isSorted(sort sorter, items []item) bool {
	for i := 1; i < len(items); i++ {
		if sort.less(items[i], items[i-1]) {
			return false
		}
	}
	return true
}

type secondsort struct {
	sorter
}

func testSorter2(base sorter, items []item) {
	show(items)
	isSorted(base, items)
}
