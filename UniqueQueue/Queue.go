package queue

type Queue interface {
	Push(key interface{})
	Pop() interface{}
	Contains(key interface{}) bool
	Len() int
	Keys() []interface{}
}

type QueueP struct {
	size int
	data []interface{}
}

func New(size int) *QueueP {
	return &QueueP{
		size: size,
	}
}

// Push values
func (q *QueueP) Push(n interface{}) *QueueP {
	if q.Len() < q.size {
		q.data = append(q.data, n)
	} else {
		q.Pop()
		q.Push(n)
	}
	return q
}

//Keys, Show current values with fixed size
func (q *QueueP) Keys() []interface{} {
	return q.data
}

//Check the queue logic condition if it was fullfilled or no
func (q *QueueP) Contains(key interface{}) bool {
	cond := false
	for i := 0; i < q.Len(); i++ {
		if q.data[i] == key {
			cond = true
		}
	}
	return cond
}

//Len
func (q *QueueP) Len() int {
	return len(q.data)
}

//Pop
func (q *QueueP) Pop() interface{} {
	if len(q.data) == 0 {
		return 0
	}
	element := q.data[0]
	q.data = q.data[1:]
	return element
}

