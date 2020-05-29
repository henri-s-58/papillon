package partition

type Partition interface {
	Val() int
}

type I struct {
	val int
}

func (i *I) Val() int {
	if i == nil {
		return 0
	}
	return i.val
}
