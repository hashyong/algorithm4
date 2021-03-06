# 学习笔记

> 基于
> 1. Sedgewick巨著 《算法》（第四版）
> 2. CSAPP
> 3. 操作系统导论
> 4. DDIA
> 5. 重构
> 6. golang/cpp
> 7. 数据结构与算法(邓俊辉版本)

## 简介
"学而时习之，不亦说乎?有朋自远方来,不亦乐乎?人不知而不愠,不亦君子乎?" 
> 出自 《论语》：学而篇

"温故而知新，可以为师矣"
> 出自 《论语》：为政篇

1. 学习算法用论语的这两句话就可概括

    工作多年，数据结构和算法知识虽早已烂熟于心，但毕竟是基础，
时常总结回顾总有新的收获。同时觉得知识体系过于零散，遂在这里将其汇总整理，以便时常翻阅
   
2. Quote by Linus Torvalds: “Talk is cheap. Show me the code.”
  掌握算法原理后，自己不写一遍能叫学会了吗？书中涉及算法都会有golang版本实现（有精力再搞一版cpp的，毕竟老本行了）
 
3. golang 标准库源码自然是涉及到各种算法以及最佳实践，源码之下，了无秘密，学习优秀的代码如何写，才能持续进步，doc下后续会有源码解析相关文章

4. 兼顾整理 操作系统，网络，数据库，项目，开发语言等基础知识, 便于自己查阅

备注：
1. 示例代码用golang编写
2. go相关工具在这里也记录下，舒服的一批
  - json2go https://mholt.github.io/json-to-go/
  - fast parse json https://github.com/tidwall/gjson
  - cli 工具
    - https://github.com/spf13/cobra
    - https://github.com/urfave/cli
    - https://golang.org/pkg/flag/
  
## 使用
1. git clone https://github.com/hashyong/algorithm4.git
2. go get -v  golang.org/x/tools/cmd/godoc
3. godoc -http=":12345"
4. http://localhost:12345/pkg/algorithm4

[![Go Reference](https://pkg.go.dev/badge/github.com/hashyong/algorithm4.svg)](https://pkg.go.dev/github.com/hashyong/algorithm4)

## 仓库结构
- [compiler](compiler)
  - 为什么写这个？
    - 编译器是理论和实践完美结合，优美的一批，值得一学
    - 学习分为两个小部分，原理+实践
  - 编译原理
  - 自制编译器
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
- [find](search)
  - 查找相关算法
    - [无序链表](search)
    - [基于有序数组的二叉查找版本](search)
    - [二叉查找树(BST)](search)
    - [平衡查找树](search)
      - [伸展树](search)
      - [B-树](search)
      - [2-3查找树](search)
      - [左倾红黑树](search)
        - 图解系列
          - [图解红黑树插入](https://github.com/hashyong/drawio/blob/main/rbtree.jpg)
          - [图解红黑树删除](https://github.com/hashyong/drawio/blob/main/rbtree.jpg)
      - [正常红黑树](search)
    - [散列表](search)
      - [基于拉链法散列表](search)
      - [基于线性探测法散列表](search)
    - [跳跃链表](search)
      - 本质为基于概率来保证查找性能
      - 相当于多级索引，加快查找速度
- [graph](graph)
  - 图相关算法
    - [无向图](graph)
      - [深度优先搜索](graph)
      - [寻找路径](graph)
      - [广度优先搜索](graph)
      - [连通分量](graph)
      - [符号图](graph)
    - [有向图](graph)
      - [数据类型](graph)
      - [可达性](graph)
      - [环&有向无环图](graph)
      - [强连通性](graph)
        - [拓扑排序](graph)
        - [Kosaraju算法](graph)
        - [传递闭包](graph)
    - [最小生成树](graph)
      - [加权无向图](graph)
      - [Prim 算法](graph)
      - [Kruskal 算法](graph)
    - [最短路径](graph)
      - [加权有向图](graph)
      - [Dijkstra算法](graph)
      - [无环加权图最短路径](graph)
      - [加权有向图最短路径](graph)
- [string](string)
  - 字符串相关算法
    - [字符串排序](string)
      - [键索引计数法](string)
      - [低位优先](string)
      - [高位优先](string)
      - [三向字符串快速排序](string)
    - [单词查找树](string)
    - [字符串查找](string)
      - [KMP](string)
      - [BM](string)
      - [RK](string)
    - [正则表达式](string)
    - [字符串压缩](string)
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