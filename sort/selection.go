package sort

import "fmt"

type selection struct {
	sorter1
}

func (s *selection) sort(items []item) {

}

func (s *selection) less(a1 item, a2 item) bool {
	fmt.Println("no imp less")
	return false
}
