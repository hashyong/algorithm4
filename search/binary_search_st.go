package search

// 为了key之间可以比较
type BSSTI interface {
	Compare(bssti BSSTI) int
}

// 二分查找， 基于有序数组
type BinarySearchST struct {
	keys   []BSSTI
	values []interface{}
	N      int
}

func InitBinarySearchST(capcity int) *BinarySearchST {
	return &BinarySearchST{
		keys:   make([]BSSTI, capcity, capcity),
		values: make([]interface{}, capcity, capcity),
		N:      0,
	}
}

func (b *BinarySearchST) Size() int {
	return b.N
}

func (b *BinarySearchST) IsEmpty() bool {
	return b.N == 0
}

// 返回小于key的数量
// 基于迭代的二分查找
func (b *BinarySearchST) rank(key BSSTI) int {
	lo := 0
	hi := b.N - 1

	for lo <= hi {
		mid := lo + (hi-lo)/2
		res := key.Compare(b.keys[mid])
		if res == 0 {
			return mid
		}

		if res < 0 {
			hi = mid - 1
		}

		if res > 0 {
			lo = mid + 1
		}
	}
	return lo
}

func (b *BinarySearchST) Get(key BSSTI) interface{} {
	if b.IsEmpty() {
		return nil
	}

	i := b.rank(key)
	if i < b.N && key.Compare(b.keys[i]) == 0 {
		return b.values[i]
	}

	return nil
}

func (b *BinarySearchST) Set(key BSSTI, value interface{}) {
	i := b.rank(key)

	// 如果已经存在则更新
	if i < b.N && key == b.keys[i] {
		b.values[i] = value
		return
	}

	// 否则要进行插入
	// 先将大于key的数向右移动一个
	for j := b.N; j > i; j-- {
		b.keys[j] = b.keys[j-1]
		b.values[j] = b.values[j-1]
	}
	b.keys[i] = key
	b.values[i] = value
	b.N++
}

func (b *BinarySearchST) Del(key BSSTI) {
	if b.IsEmpty() {
		return
	}

	i := b.rank(key)

	// 如果已经存在则更新
	if i < b.N && key == b.keys[i] {
		for j := i; j < b.N; j++ {
			b.keys[j] = b.keys[j+1]
			b.values[j] = b.values[j+1]
		}
		b.N--
	}

	// 如果不存在直接返回即可， 无用处理
}

func (b *BinarySearchST) Min() interface{} {
	return b.keys[0]
}

func (b *BinarySearchST) Max() interface{} {
	return b.keys[b.N-1]
}

func (b *BinarySearchST) Select(k int) interface{} {
	return b.keys[k]
}
