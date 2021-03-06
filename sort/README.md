# 排序算法原理简介

1. [选择排序](#选择排序)
2. [插入排序](#插入排序)
3. [希尔排序](#希尔排序)
4. [归并排序](#归并排序)
5. [快速排序](#快速排序)
6. [堆排序](#堆排序)

## 选择排序
- 简介： 一种最简单的排序算法
  - 首先，找到数组中最小的那个元素，
  - 其次，将它和数组的第 一个元素交换位置(如果第一个元素就是最小元素那么它就和自己交换)。
  - 再次，在剩下的元素中 找到最小的元素，将它与数组的第二个元素交换位置。
  - 如此往复，直到将整个数组排序。这种方法 叫做选择排序，因为它在不断地选择剩余元素之中的最小者。
- 特点
  - 运行时间和输入无关。为了找出最小的元素而扫描一遍数组并不能为下一遍扫描提供什么信息
  - 数据移动是最少的。每次交换都会改变两个数组元素的值，因此选择排序用了 N 次交换——交 换次数和数组的大小是线性关系
- [code](sort.go#L36)
## 插入排序
- 简介
  - 通常人们整理桥牌的方法是一张一张的来，将每一张牌插入到其他已经有序的牌中的适当位置。 在计算机的实现中，为了给要插入的元素腾出空间，我们需要将其余所有元素在插入之前都向右移 动一位。这种算法叫做插入排序
  - 与选择排序一样，当前索引左边的所有元素都是有序的，但它们的最终位置还不确定，为了给 更小的元素腾出空间，它们可能会被移动。但是当索引到达数组的右端时，数组排序就完成了。
  - 和选择排序不同的是，插入排序所需的时间取决于输入中元素的初始顺序。例如，对一个很大 且其中的元素已经有序(或接近有序)的数组进行排序将会比对随机顺序的数组或是逆序数组进行 排序要快得多。
- [code](sort.go#L57)
## 希尔排序
- 简介
  - 为了展示初级排序算法性质的价值，接下来我们将学习一种基于插入排序的快速的排序算法。
  - 对于大规模乱序数组插入排序很慢，因为它只会交换相邻的元素，因此元素只能一点一点地从数组 的一端移动到另一端。
    例如，如果主键最小的元素正好在数组的尽头，要将它挪到正确的位置就需 要 N-1 次移动。
    希尔排序为了加快速度简单地改进了插入排序，交换不相邻的元素以对数组的局部 进行排序，并最终用插入排序将局部有序的数组排序。
  - 希尔排序的思想是使数组中任意间隔为 h 的元素都是有序的。这样的数组被称为 h 有序数组。
   换句话说，一个 h 有序数组就是 h 个互相独立的有序数组编织在一起组成的一个数组
   在进行排序时，如果 h 很大，我们就能将元素移动到很远的地方，为实现更小的 h 有序创造方便。用
   这种方式，对于任意以 1 结尾的 h 序列，我们都能够将数组排序。这就是希尔排序。
   使用了序列 1/2(3 -1)，从 N/3 开始递减至 1。我们把这个序列称为递增序列。
- [code](sort.go#L80)
## 归并排序
- 简介
  - 我们所讨论的算法都基于归并这个简单的操作，即将两个有序的数组归并成一个更大 的有序数组。
    很快人们就根据这个操作发明了一种简单的递归排序算法:归并排序。
    要将一个数组 排序，可以先(递归地)将它分成两半分别排序，然后将结果归并起来。
    你将会看到，归并排序最 吸引人的性质是它能够保证将任意长度为 N 的数组排序所需时间和 NlogN 成正比;
    它的主要缺点 则是它所需的额外空间和 N 成正比
  - 自顶向下的归并排序
    - 简介
      - 基于原地归并的抽象实现了另一种递归归并，这也是应用高效算法设计中分治思想的 最典型的一个例子。
        这段递归代码是归纳证明算法能够正确地将数组排序的基础:如果它能将两个 子数组排序，它就能够通过归并两个子数组来将整个数组排序。
      - 优化：对小规模子数组使用插入排序
  - 自底向上的归并排序 
    - 简介
      - 实现归并排序的另一种方法是先归并那些微型数组，然后再 成对归并得到的子数组，如此这般，直到我们将整 个数组归并在一起。
        这种实现方法比标准递归方法 所需要的代码量更少。首先我们进行的是两两归并
        (把每个元素想象成一个大小为 1 的数组)，然后 是四四归并(将两个大小为 2 的数组归并成一个有 4 个元素的数组)，
        然后是八八的归并，一直下去。 
        在每一轮归并中，最后一次归并的第二个子数组可 能比第一个子数组要小(但这对 merge() 方法不是 问题)，如果不是的话所有的归并中两个数组大小 都应该一样，而在下一轮中子数组的大小会翻倍。
- [自底向上代码](sort.go#L124)
- [自顶向下代码](sort.go#L146)
## 快速排序
- 简介
  - 快速排序是一种分治的排序算法。
    - 它将一个数组分成两个子数组，将两部分独立地排序。
    - 快速排序和归并排序是互补的:归并排序将数组分成两个子数组分别排序，并将有序的子数组归并以将整个 数组排序;
      而快速排序将数组排序的方式则是当两个子数组都有序时整个数组也就自然有序了。
      在第一种情况中，递归调用发生在处理整个数组之前;
      在第二种情况中，递归调用发生在处理整个数组之后。 
    - 在归并排序中，一个数组被等分为两半;在快速排序中，切分(partition)的位置取决于数组的内容。
  - 基本实现： 二向切分
    - 快速排序递归地将子数组 a[lo..hi] 排序，先用 partition() 方法将 a[j] 放到一个合适位置，然
      后再用递归调用将其他位置的元素排序。
      该方法的关键在于切分，这个过程使得数组满足下面三个条件:
      -  对于某个 j，a[j] 已经排定;
      -  a[lo] 到 a[j-1] 中的所有元素都不大于 a[j];
      -  a[j+1] 到 a[hi] 中的所有元素都不小于 a[j]。
      
      我们就是通过递归地调用切分来排序的。
    
  - 基于重复元素优化： 三向切分
    - 实际应用中经常会出现含有大量重复元素的数组排序，一个元素全部重复的子数组就不需要继续排序了，但我们的算法还会继续将它切分 为更小的数组。
      在有大量重复元素的情况下，快速排序的递归性会使元素全部重复的子数组经常出现，这就有很大的改进潜力，将当前实现的线性对数级的性能提高到线性级别。
    - 原理
      - 一个简单的想法是将数组切分为三部分，分别对应小于、等于和大于切分元素的数组元素
      - Dijkstra 的解法如“三向切分的快速排序”中极为简洁的切分代码所示。
        它从左到右遍历数组 一次，
        维护一个指针 lt 使得 a[lo..lt-1] 中的元素都小于 v，
        一个指针 gt 使得 a[gt+1..hi] 中 的元素都大于 v，
        一个指针 i 使得 a[lt..i-1] 中的元素都等于 v，
        a[i..gt] 中的元素都还未确定， 

        一开始 i 和 lo 相等，我们使用 Comparable 接口(而非 less())对 a[i] 进行三 向比较来直接处理以下情况
        - a[i]小于v，将a[lt]和a[i]交换，将lt和i加一;
        -  a[i] 大于 v，将 a[gt] 和 a[i] 交换，将 gt 减一;
        - a[i]等于v，将i加一。
        
        这些操作都会保证数组元素不变且缩小 gt-i 的值(这样循环才会结束)。另外，除非和切分
        元素相等，其他元素都会被交换。
## 堆排序
- 简介
  - 使用二叉堆来实现
  - 堆的上浮或者下沉
  - 排序需先初始化来构造堆， 每次将新元素插入到堆的末尾，然后上浮即可
  - 排序的时候 每次删除堆顶元素， 最小值， 再将最后一个元素和堆顶元素交换，再下沉，如此往复即可
- 标准库实现非常优雅，可以直接参考[PQ](https://golang.org/pkg/container/heap/)
- [删除指定元素原理](http://www.mathcs.emory.edu/~cheung/Courses/171/Syllabus/9-BinTree/heap-delete.html)
- [堆排序改进：先下沉后上浮](https://zhuanlan.zhihu.com/p/28593993)    