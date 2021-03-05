package sort

import (
	"algorithm4/base"
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
	ret.DisplayData = append(ret.DisplayData, data.Copy().(ISort))

	end := data.Len()
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

// 希尔排序
// 插入排序， 当最小值在最后的时候 交换次数过多，N -1，每次可以多跳几个
func ShellSort(data ISort) Display {
	var ret Display
	ret.DisplayData = append(ret.DisplayData, data.Copy().(ISort))

	// 本质上是一个递增序列+插入排序, 加快向后移动的速度
	// 两种实现
	// 1. 每个数组独立排序
	// 2. 每个数组独立，将其交换至比其大的元素前边就行
	N := data.Len()
	h := 1

	// 递增序列选择 1, 4, 13, 40 。。。
	for {
		if h >= N/3 {
			break
		}
		h = h*3 + 1
	}

	for {
		if h < 1 {
			break
		}
		for i := h; i < N; i++ {
			for j := i; j >= h && data.Less(j, j-h); j -= h {
				data.Swap(j, j-h)
				ret.DisplayData = append(ret.DisplayData, data.Copy().(ISort))
			}
		}
		h = h / 3
	}

	return ret
}

// 归并排序

// 快排

// 队排序

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

func Sort(name string) {
	fmt.Println(name, " begin")
	in := make(base.Items, base.Inputs.Len())
	copy(in, base.Inputs)

	switch name {
	case "insert":
		Show(InsertSort(in))
	case "bubble":
		Show(BubbleSort(in))
	case "shell":
		Show(ShellSort(in))
	default:
	}
	fmt.Println(name, " end")
}
