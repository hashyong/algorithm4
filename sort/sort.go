package sort

import (
	"algorithm4/base"
	_ "container/heap"
	"fmt"
	_ "sort"
	"time"
)

// 排序通用接口
type ISort interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)

	// 返回当前数据状态
	Copy() interface{}
	// 归并排序使用, 后边这里再想想咋个优化，因为不需要临时空间的排序都不需要这个
	Set(int, interface{})
	Get(int) interface{}
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

func Min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}

// 归并排序
// 自下而上
func MergeBUSort(data ISort) Display {
	var ret Display
	ret.DisplayData = append(ret.DisplayData, data.Copy().(ISort))

	aux := data.Copy().(ISort)
	for sz := 1; sz < data.Len(); sz = sz + sz {
		for lo := 0; lo < data.Len()-sz; lo += sz + sz {
			merge(data, aux, lo, lo+sz-1, Min(lo+sz+sz-1, data.Len()-1))
		}
	}
	return ret
}

// 自顶向下
func MergeSort(data ISort) Display {
	var ret Display
	ret.DisplayData = append(ret.DisplayData, data.Copy().(ISort))

	sortMerge(data, data.Copy().(ISort), 0, data.Len()-1)
	return ret
}

func sortMerge(data ISort, aux ISort, lo, hi int) {
	if hi <= lo {
		return
	}

	mid := lo + (hi-lo)/2
	sortMerge(data, aux, lo, mid)
	sortMerge(data, aux, mid+1, hi)
	merge(data, aux, lo, mid, hi)
}

// 辅助方法 merge
func merge(data ISort, aux ISort, lo, mid, hi int) {
	i := lo
	j := mid + 1

	aux = data.Copy().(ISort)

	// i     j
	// 3 4 5 1 9 10
	// 1
	for k := lo; k <= hi; k++ {
		if i > mid {
			data.Set(k, aux.Get(j))
			j++
			continue
		}

		if j > hi {
			data.Set(k, aux.Get(i))
			i++
			continue
		}

		if aux.Less(j, i) {
			data.Set(k, aux.Get(j))
			j++
		} else {
			data.Set(k, aux.Get(i))
			i++
		}
	}
}

// 快排 双向切分
func QuickSort(data ISort) Display {
	var ret Display
	ret.DisplayData = append(ret.DisplayData, data.Copy().(ISort))

	quickSort(data, 0, data.Len()-1)
	return ret
}

// 快排 三向切分
func Quick3WaySort(data ISort) Display {
	var ret Display
	ret.DisplayData = append(ret.DisplayData, data.Copy().(ISort))

	quickSort(data, 0, data.Len()-1)
	return ret
}

// 子数组， 依赖 compare方法，能返回三个状态，大于，小于，或者等于, 目前未实现, 后续再实现
//Dijkstra 的解法如“三向切分的快速排序”中极为简洁的切分代码所示。
//它从左到右遍历数组 一次，维护一个指针 lt 使得 a[lo..lt-1] 中的元素都小于 v，
//一个指针 gt 使得 a[gt+1..hi] 中 的元素都大于 v，
//一个指针 i 使得 a[lt..i-1] 中的元素都等于 v，a[i..gt] 中的元素都还未确定，
//j如图 2.3.4 所示。一开始 i 和 lo 相等，我们使用 Comparable 接口(而非 less())对 a[i] 进行三 向比较来直接处理以下情况:
// a[i]小于v，将a[lt]和a[i]交换，将lt和i加一;
// a[i] 大于 v，将 a[gt] 和 a[i] 交换，将 gt 减一;
// a[i]等于v，将i加一。
//这些操作都会保证数组元素不变且缩小 gt-i 的值(这样循环才会结束)。另外，除非和切分
//元素相等，其他元素都会被交换。
func quick3Sort(data ISort, lo, hi int) {
	if hi <= lo {
		return
	}

	// lt := lo
	i := lo + 1
	gt := hi

	for i <= gt {

	}

}

func quickSort(data ISort, lo int, hi int) {
	if hi <= lo {
		return
	}

	j := partition(data, lo, hi)
	quickSort(data, lo, j-1)
	quickSort(data, j+1, hi)
}

func partition(data ISort, lo int, hi int) int {
	// 左右扫描指针
	i := lo
	j := hi + 1

	for {
		i++
		for data.Less(i, lo) {
			if i == hi {
				break
			}
			i++
		}

		j--
		for data.Less(lo, j) {
			if j == lo {
				break
			}
			j--
		}

		// 此时，data[i] 大于data lo， data j 小于data lo
		if i >= j {
			break
		}
		data.Swap(i, j)
	}

	data.Swap(lo, j)
	return j
}

// 堆排序
// go 标准库已经有实现，可以参考下
// container/heap/example_pq_test.go
func HeapSort(data ISort) Display {
	var ret Display
	ret.DisplayData = append(ret.DisplayData, data.Copy().(ISort))

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
	case "merge":
		Show(MergeSort(in))
	case "mergeBU":
		Show(MergeBUSort(in))
	case "quick":
		Show(QuickSort(in))
	case "heap":
		Show(HeapSort(in))
	}
	fmt.Println(name, " end")
}
