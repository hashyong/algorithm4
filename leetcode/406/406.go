package leetcode

import (
	"sort"
)

// 406. 根据身高重建队列
// 假设有打乱顺序的一群人站成一个队列。
// 每个人由一个整数对(h, k)表示，其中h是这个人的身高，k是排在这个人前面且身高大于或等于h的人数。
// 编写一个算法来重建这个队列。
//
// 注意：
// 总人数少于1100人。
//
// 示例
//
// 输入:
// [[7,0], [4,4], [7,1], [5,0], [6,1], [5,2]]
//
// 输出:
// [[5,0], [7,0], [5,2], [6,1], [4,4], [7,1]]

// [7,0], [7,1], [6,1], [5,0], [5,2], [4,4]
// 再一个一个插入。
// [7,0]
// [7,0], [7,1]
// [7,0], [6,1], [7,1]
// [5,0], [7,0], [6,1], [7,1]
// [5,0], [7,0], [5,2], [6,1], [7,1]
// [5,0], [7,0], [5,2], [6,1], [4,4], [7,1]

type sortable [][]int

func (s sortable) Len() int {
	return len(s)
}

func (s sortable) Less(i, j int) bool {
	if s[i][0] == s[j][0] {
		return s[i][1] <= s[j][1]
	}
	return s[i][0] >= s[j][0]
}

func (s sortable) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func insert(people [][]int, index int, value int) [][]int {
	copy(people[index+1:], people[index:])
	people[index] = make([]int, 2)
	people[index][0] = value
	people[index][1] = index
	return people
}

func reconstructQueue(people [][]int) [][]int {
	sort.Sort(sortable(people))

	res := make([][]int, len(people))
	for _, value := range people {
		res = insert(res, value[1], value[0])
	}

	return res
}
