package queue

import (
	"container/list"
	"sync"
)

func New() *Queue {
	return &Queue{
		list: list.New(),
	}
}

type Queue struct {
	list *list.List
	lock sync.Mutex
}

func (q *Queue) PushBack(item ...interface{}) {
	q.lock.Lock()
	for _, v := range item {
		q.list.PushBack(v)
	}
	q.lock.Unlock()
}

func (q *Queue) PushFront(item ...interface{}) {
	q.lock.Lock()
	for _, v := range item {
		q.list.PushFront(v)
	}
	q.lock.Unlock()
}

func (q *Queue) Back() interface{} {
	q.lock.Lock()
	defer q.lock.Unlock()
	v := q.list.Back()
	if v != nil {
		return q.list.Remove(v)
	}
	return nil
}

func (q *Queue) Front() interface{} {
	q.lock.Lock()
	defer q.lock.Unlock()
	v := q.list.Front()
	if v != nil {
		return q.list.Remove(v)
	}
	return nil
}

func (q *Queue) Len() int {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.list.Len()
}

func (q *Queue) Clear() {
	q.lock.Lock()
	q.list.Init()
	q.lock.Unlock()
}
