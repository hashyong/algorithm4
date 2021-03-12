package search

// 二分查找， 基于有序数组
type BinarySearchST struct {
	keys   []interface{}
	values []interface{}
	N      int
}

func InitBinarySearchST(capcity int) *BinarySearchST {
	return &BinarySearchST{
		keys:   make([]interface{}, capcity, capcity),
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
func (b *BinarySearchST) rank(key interface{}) int {
	return 0
}

func (b *BinarySearchST) Get(key interface{}) interface{} {
	if b.IsEmpty() {
		return nil
	}

	i := b.rank(key)
	if i < b.N && key == b.keys[i] {
		return b.values[i]
	}

	return nil
}

func (b *BinarySearchST) Set(key interface{}, value interface{}) {
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

func (b *BinarySearchST) Del(key interface{}) {
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
