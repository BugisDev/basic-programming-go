package slicer

type Int []int

func (sl *Int) Add(v int) {
	*sl = append(*sl, v)
}

func (sl *Int) IndexOf(v int) int {
	for i, value := range *sl {
		if value == v {
			return i
		}
	}
	return -1
}

func (sl *Int) IsExist(v int) bool {
	return sl.IndexOf(v) != -1
}

func (sl *Int) Len() int {
	return len(*sl)
}

func (sl *Int) Merge(another Int) {
	*sl = append(*sl, another...)
}
