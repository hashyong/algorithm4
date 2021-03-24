# 查找算法

## 最最最基础的无序链表版本 
- [SeqSearchST](search.go)

## 基于有序数组的二叉查找版本
- [BinarySearchST](binary_search_st.go)

## BST
- [BST](binary_search_tree.go)
  - 支持以下接口 
    - get, put, size, range, select, floor, delmin, delmax, del

## 左倾红黑树  
- [LeftRedBlackTree](left_rb_tree.go)
- 图解插入流程
  
![图解插入](https://github.com/hashyong/drawio/blob/main/rbtree.jpg)

## hash表
- 过于简单， 此处不展开
- 解决冲突分为两种方法
  1. 基于拉链法， 冲突之后往链表上挂数据就行， 假如太多， 建议转为红黑树
  2. 基于线性探测法， 往后找，有空的直接插入即可