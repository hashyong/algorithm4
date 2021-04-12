package graph

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type UF struct {
	// 分量id， 以触点做为索引
	id []int
	// 分量数量
	count int

	// 根节点对应分量大小
	// 保证合并后 树的深度不会过深
	// 小树合并到大树根节点去
	sz []int
}

func NewUF() *UF {
	return &UF{}
}

func (u *UF) init(n int) {
	u.count = n
	u.id = make([]int, n)
	for i := 0; i < n; i++ {
		u.id[i] = i
	}
}

func (u *UF) initWeight(n int) {
	u.count = n
	u.id = make([]int, n)
	u.sz = make([]int, n)
	for i := 0; i < n; i++ {
		u.id[i] = i
		u.sz[i] = 1
	}
}

func (u *UF) Read(name string) {
	// https://mholt.github.io/json-to-go/
	type Autogenerated struct {
		V    int `json:"v"`
		Edge []struct {
			S int `json:"s"`
			E int `json:"e"`
		} `json:"edge"`
	}

	// 读取标准输入数据
	content, err := ioutil.ReadFile(name)
	if err != nil {
		return
	}

	// json 转换
	var d Autogenerated
	err = json.Unmarshal(content, &d)
	if err != nil {
		return
	}

	// 构造数据
	u.init(d.V)
	for _, data := range d.Edge {
		if u.Connected(data.E, data.S) {
			continue
		}

		u.Union(data.S, data.E)
		fmt.Println(data.S, " ", data.E)
	}
	fmt.Println(u.count, " components")
}

func (u *UF) Count() int {
	return u.count
}

func (u *UF) Connected(p, q int) bool {
	return u.Find(p) == u.Find(q)
}

func (u *UF) Union(p, q int) {
	fmt.Println(p, q)
}

func (u *UF) Find(p int) int {
	return p
}

func (u *UF) QFFind(p int) int {
	return u.id[p]
}

func (u *UF) QFUnion(p, q int) {
	pID := u.Find(p)
	qID := u.Find(q)

	if pID == qID {
		return
	}

	for i := 0; i < len(u.id); i++ {
		if u.id[i] == pID {
			u.id[i] = qID
		}
	}

	u.count--
}

// QUSubFind 返回对应根节点
func (u UF) QUSubFind(p int) int {
	for p != u.id[p] {
		p = u.id[p]
	}
	return p
}

func (u *UF) QUUnion(p, q int) {
	pRoot := u.QUSubFind(p)
	qRoot := u.QUSubFind(q)

	if pRoot == qRoot {
		return
	}

	u.id[pRoot] = qRoot
	u.count--
}

func (u *UF) QUWeightUnion(p, q int) {
	pRoot := u.QUSubFind(p)
	qRoot := u.QUSubFind(q)

	if pRoot == qRoot {
		return
	}

	// 妙的一批
	if u.sz[pRoot] < u.sz[qRoot] {
		u.id[pRoot] = qRoot
		u.sz[qRoot] += u.sz[qRoot]
	} else {
		u.id[qRoot] = pRoot
		u.sz[qRoot] += u.sz[pRoot]
	}

	u.count--
}
