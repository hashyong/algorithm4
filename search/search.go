package search

import "strings"

type LinkNode struct {
	Key  BSSTI
	Data interface{}
	Next *LinkNode
	Prev *LinkNode
}

// 无序链表
type SeqSearchST struct {
	Head *LinkNode
	Tail *LinkNode
	size int
}

func InitNode(key BSSTI, data interface{}) *SeqSearchST {
	res := &LinkNode{
		Key:  key,
		Data: data,
		Next: nil,
		Prev: nil,
	}

	table := &SeqSearchST{
		Head: res,
		Tail: res,
		size: 1,
	}
	return table
}

// 无序链表, 查找只能挨个遍历
func (c *SeqSearchST) Get(key interface{}) interface{} {
	node := c.Head
	for node != nil {
		if node.Key == key {
			return node.Data
		}

		node = (*node).Next
	}

	return nil
}

// 无序链表，直接将元素设置到链表结尾即可
func (c *SeqSearchST) Set(key BSSTI, value interface{}) {
	node := &LinkNode{
		Key:  key,
		Data: value,
		Next: nil,
		Prev: nil,
	}

	node.Prev = c.Tail
	c.Tail.Next = node
	c.Tail = node
	c.size++
}

// 返回无序列表size
func (c *SeqSearchST) Size() int {
	return c.size
}

// 删除某个指定的key
func (c *SeqSearchST) Del(key interface{}) {
	node := c.Head
	for node != nil {
		if node.Key == key {
			node.Prev.Next = node.Next
			return
		}

		node = node.Next
	}
}

// 不支持范围查找, 原始链表无序
// 要支持，可以先排序 之后再返回某个范围的值

type Key struct {
	k string
}

func (i *Key) Compare(j BSSTI) int {
	return strings.Compare(i.k, j.(*Key).k)
}
