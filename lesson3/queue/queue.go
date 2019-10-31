// Package queue provides
package queue

import "fmt"

type Queue []int

func NewQueue(size int) Queue {
	return make([]int, 0, size)
}

func (q *Queue) Push(v int) {
	(*q) = append((*q), v)
}

func (q *Queue) Pop() int {
	result := (*q)[0]
	(*q) = (*q)[1:]
	return result
}

func (q Queue) Print() {
	fmt.Println(q)
}
