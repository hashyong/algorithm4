package leetcode

import (
	"fmt"
	"testing"
)

// [[7,0], [4,4], [7,1], [5,0], [6,1], [5,2]]
func Test406(t *testing.T) {
	f := [][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}
	fmt.Println(f)
	fmt.Println("res:", reconstructQueue(f))
}
