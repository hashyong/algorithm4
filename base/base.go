package base

type Item struct {
	id int
}

type Items []Item

func (r Items) Len() int {
	return len(r)
}

func (r Items) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r Items) Less(i, j int) bool {
	return r[i].id < r[j].id
}

func (r Items) Copy() interface{} {
	tmp := make(Items, r.Len())
	copy(tmp, r)
	return tmp
}

var Inputs = Items{{74}, {59}, {23}, {12}, {-784}, {9845}}
