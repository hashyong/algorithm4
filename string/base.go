package string

// Alphabet 字母表
type Alphabet struct{}

func NewAlphabet() *Alphabet {
	return &Alphabet{}
}

// Alphabet 根据s的中字符常见一张新的字母表
func (a *Alphabet) Alphabet(s string) {

}

// 获取字母表中索引位置的字符
func (a *Alphabet) toChar(index int) byte {
	return '0'
}

// 获取b的索引
func (a *Alphabet) toIndex(b byte) int {
	return -1
}

// b 是否在字母表中
func (a *Alphabet) contains(b byte) bool {
	return false
}

// R 基数, 字母表中的字符数量
func (a *Alphabet) R() int {
	return 0
}

// 表示一个索引需要的比特数
func (a *Alphabet) lgR() int {
	return 0
}

// 将s转换为R进制的整数
func (a *Alphabet) toIndices(s string) []int {
	return nil
}

// 将R进制的整数转化为 基于该字母表的字符串
func (a *Alphabet) toChars(indices []int) string {
	return ""
}
