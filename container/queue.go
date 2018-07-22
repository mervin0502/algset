package container

//Queue define the queue structure
type Queue struct {
	head *Item
	end  *Item
}

//NewQueue create a new queue
func NewQueue() *Queue {
	q := &Queue{nil, nil}
	return q
}

//Push append a item
func (q *Queue) Push(data interface{}) {
	n := &Item{data: data, next: nil}

	if q.end == nil {
		q.head = n
		q.end = n
	} else {
		q.end.next = n
		q.end = n
	}

	return
}

//Pop get the queue head and remove it
func (q *Queue) Pop() (interface{}, bool) {
	if q.head == nil {
		return nil, false
	}

	data := q.head.data
	q.head = q.head.next
	if q.head == nil {
		q.end = nil
	}

	return data, true
}

//Empty return true if the queue is empty
func (q *Queue) Empty() bool {
	if q.head == nil && q.end == nil {
		return true
	}
	return false
}
