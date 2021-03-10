package main

import (
	msort "algorithm4/sort"
)

func main() {
	var sortList []string
	sortList = append(sortList, "insert")
	sortList = append(sortList, "bubble")
	sortList = append(sortList, "shell")
	sortList = append(sortList, "merge")
	sortList = append(sortList, "mergeBU")
	sortList = append(sortList, "quick")
	for _, s := range sortList {
		msort.Sort(s)
	}
}
