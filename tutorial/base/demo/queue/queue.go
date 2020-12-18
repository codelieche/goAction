package queue

// Queue只支持int类型
type Queue []int

// Queue2支持任何类型
type Queue2 []interface{}

func (q *Queue) Push(v int) {
	*q = append(*q, v)
}

func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

// Queue2

func (q *Queue2) Push(v interface{}) {
	// 如果只想Push int，那么把v的类型设置为int即可
	*q = append(*q, v)
}

func (q *Queue2) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	// 如果只想Pop出int
	//return head.(int)
	return head
}

func (q *Queue2) IsEmpty() bool {
	return len(*q) == 0
}
