package sort

import "fmt"

type ISort interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

func BubbleSort(data ISort) {
	if data.Len() <= 1 {
		return
	}

	end := data.Len()
	for i := 0; i < end; end-- {
		fmt.Println(data)
		for j := i + 1; j < end; j++ {
			if data.Less(j, j-1) {
				data.Swap(j, j-1)
			}
		}
	}
}

// IsSorted reports whether data is sorted.
func IsSorted(data ISort) bool {
	n := data.Len()
	for i := n - 1; i > 0; i-- {
		if data.Less(i, i-1) {
			return false
		}
	}
	return true
}
