package string

import "errors"

// Alphabet 字母表
type Alphabet struct {
	// the characters in the alphabet
	alphabet []byte
	// indices
	inverse []int
	// the radix of the alphabet
	r int
}

func NewAlphabet(t string) (*Alphabet, error) {
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
		return ret, errors.New("default")
	}

	return ret.Alphabet(param)
}

// Alphabet 根据s的中字符常见一张新的字母表
// Initializes a new alphabet from the given set of characters
func (a *Alphabet) Alphabet(alpha string) (*Alphabet, error) {
	// check that alphabet contains no duplicate chars
	unicode := make(map[byte]bool, 65535)
	for i := 0; i < len(alpha); i++ {
		if unicode[alpha[i]] {
			return a, errors.New("repeat")
		}
		unicode[alpha[i]] = true
	}

	a.alphabet = []byte(alpha)
	a.r = len(alpha)
	a.inverse = make([]int, 65535)
	//inverse = new int[Character.MAX_VALUE];
	//for (int i = 0; i < inverse.length; i++)
	//inverse[i] = -1;
	//
	//// can't use char since R can be as big as 65,536
	//for (int c = 0; c < R; c++)
	//inverse[alphabet[c]] = c;
	//

	return a, nil
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
