package string

// Alphabet 字母表
type Alphabet struct {
	// the characters in the alphabet
	alphabet []byte
	// indices
	inverse []int
	// the radix of the alphabet
	r int
}

func NewAlphabet(t string) *Alphabet {
	ret := &Alphabet{}
	param := ""
	switch t {
	case "bin":
		param = "01"
	case "oct":
		param = "01234567"
	case "dec":
		param = "0123456789"
	default:
		return ret
	}

	return ret.Alphabet(param)
}

// Alphabet 根据s的中字符常见一张新的字母表
// Initializes a new alphabet from the given set of characters
func (a *Alphabet) Alphabet(s string) *Alphabet {

	return a
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
