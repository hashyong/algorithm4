package main

import (
	"algorithm4/base"
	msort "algorithm4/sort"
	"fmt"
)

func main() {
	fmt.Println("InsertSort begin")
	in1 := make(base.Items, base.Inputs.Len())
	copy(in1, base.Inputs)
	msort.Show(msort.InsertSort(in1))
	fmt.Println("InsertSort end")

	fmt.Println("BubbleSort begin")
	in2 := make(base.Items, base.Inputs.Len())
	copy(in2, base.Inputs)
	msort.Show(msort.BubbleSort(in2))
	fmt.Println("BubbleSort end")
}
