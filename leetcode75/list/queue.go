package list

import (
	"iter"
)

type Queue[T comparable] struct {
	elements []T
}

//type Queue []TreeNode

func (q *Queue[T]) IsEmpty() bool {
	return (*q).Size() == 0
}

func (q *Queue[T]) Add(node *T) {
	q.elements = append(q.elements, *node)
}

func (q *Queue[T]) Remove() *T {
	node := q.elements[0]
	q.elements = q.elements[1:]
	return &node
}

func (q *Queue[T]) GetLast() *T {
	return &q.elements[len(q.elements)-1]
}
func (q *Queue[T]) Size() int {
	return len(q.elements)
}

func (q *Queue[T]) AddAll(anotherQ *Queue[T]) {
	for node := range anotherQ.Iter() {
		q.Add(&node)
	}
}

func (q *Queue[T]) Iter() iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, element := range q.elements {
			if !yield(element) {
				return
			}
		}
	}
}

func NewQueue[T comparable]() *Queue[T] {
	return &Queue[T]{
		elements: []T{},
	}
}
