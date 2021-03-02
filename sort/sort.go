package sort

import (
	"fmt"
	"time"
)

// 排序通用接口
type ISort interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)

	Copy() interface{}
}

type Display struct {
	DisplayData []ISort
}

func Show(data Display) {
	for i, item := range data.DisplayData {
		fmt.Printf("Res step i%4d: %v\r", i, item)
		time.Sleep(time.Second * 1)
	}
}

// 冒泡排序
func BubbleSort(data ISort) Display {
	var ret Display
	ret.DisplayData = append(ret.DisplayData, data.Copy().(ISort))
	if data.Len() <= 1 {
		return ret
	}

	end := data.Len()
	for i := 0; i < end; end-- {
		for j := i + 1; j < end; j++ {
			if data.Less(j, j-1) {
				data.Swap(j, j-1)
				ret.DisplayData = append(ret.DisplayData, data.Copy().(ISort))
			}
		}
	}

	return ret
}

// 插入排序
func InsertSort(data ISort) Display {
	var ret Display

	end := data.Len()
	ret.DisplayData = append(ret.DisplayData, data.Copy().(ISort))

	for i := 1; i < end; i++ {
		//  i 挨个插入
		// 5 2 1
		// 5
		// 5 2 => 2 5
		// 2 1 5
		// 1 2 5
		for j := i; j > 0 && data.Less(j, j-1); j-- {
			data.Swap(j, j-1)
			ret.DisplayData = append(ret.DisplayData, data.Copy().(ISort))
		}
	}

	return ret
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
