package utils

type Queue[T any] struct {
	Elements []T
	Maxsize  int
}

func NewQueue[T any](maxsize int) Queue[T] {
	return Queue[T]{Elements: make([]T, maxsize), Maxsize: maxsize}
}

func NewQueueFromList[T any](list []T, maxsize int) Queue[T] {
	if len(list) > maxsize && maxsize != 0 {
		panic("The provided list cannot be longer than the maximum size of the Queue")
	}
	return Queue[T]{Elements: list, Maxsize: maxsize}
}

func (q *Queue[T]) Put(value T) {
	if len(q.Elements)+1 > q.Maxsize && q.Maxsize != 0 {
		panic("The Queue is full and no further element can be added")
	}
	q.Elements = append(q.Elements, value)
}

/*func (q *Queue[T]) Push(value T) {
	if len(q.Elements)+1 > q.Maxsize && q.Maxsize != 0 {
		panic("The Queue is full and no further element can be added")
	}
	q.Elements = append(q.Elements, value)
}*/

func (q *Queue[T]) Get() T {
	if q.Empty() {
		panic("The Queue is empty")
	}
	v := q.Elements[0]
	q.Elements = q.Elements[1:]
	return v
}

/*func (q *Queue[T]) Pop() T {
	if q.Empty() {
		panic("The Queue is empty")
	}
	v := q.Elements[0]
	q.Elements = q.Elements[1:]
	return v
}*/

func (q *Queue[T]) Empty() bool {
	return len(q.Elements) == 0
}
