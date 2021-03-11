# 学习笔记

> 基于
> 1. Sedgewick巨著 《算法》（第四版）

## 简介
"学而时习之，不亦说乎?有朋自远方来,不亦乐乎?人不知而不愠,不亦君子乎?" 
> 出自 《论语》：学而篇

"温故而知新，可以为师矣"
> 出自 《论语》：为政篇

1. 学习算法用论语的这两句话就可概括

    工作多年，数据结构和算法知识虽早已烂熟于心，但毕竟是基础，
时常总结回顾总有新的收获。同时觉得知识体系过于零散，遂在这里将其汇总整理，以便时常翻阅
   
2. 兼顾整理 操作系统，网络，数据库，项目，开发语言等基础知识, 便于自己查阅

备注：
1. 示例代码用golang编写

## 仓库结构
- [base](base)
  - 基础数据结构在此定义
- [sort](sort)
  - 排序相关算法
    - 初级
      1. [选择排序](sort)
      2. [插入排序](sort)
      3. [希尔排序](sort)
    - 归并排序
      1. [自顶向下](sort)
      2. [自下而上](sort)
    - 快速排序
      1. [标准二向切分](sort)
      2. [三向切分](sort)
          - 对于处理重复元素， 效果较好
    - 堆排序（优先队列）
      - 标准库实现非常优雅，可以直接参考[PQ](https://golang.org/pkg/container/heap/)
      - [删除指定元素原理](http://www.mathcs.emory.edu/~cheung/Courses/171/Syllabus/9-BinTree/heap-delete.html)
      - [堆排序改进：先下沉后上浮](https://zhuanlan.zhihu.com/p/28593993)
    - TODO：分析golang sort 源码
- [find](find)
  - 查找相关算法
- [graph](graph)
  - 图相关算法
- [string](string)
  - 字符串相关算法
- [search](search)
  - 搜索相关算法
- [leetcode](leetcode)
  - 相关题解
- [algorithm_mianshi](algorithm_mianshi)
  - 面试相关算法
- [database](database)
  - 数据库相关
- [network](network)
  - 网络相关
- [os](os)
  - 操作系统相关
- [project](project)
  - 工作相关项目
- [doc](doc)
  - 相关资料&书籍
    - [算法4 高清文字版](doc/算法（第4版）文字版.pdf)

## FAQ

- 有写的不对或不准确的地方， 欢迎发起issue, 或者邮件讨论(邮件地址见个人简介)