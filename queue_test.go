package queue

import "testing"

var (
	q  = New()
	ts = "test string"
	ti = 0
)

func BenchmarkPushBack(t *testing.B) {
	q.PushBack(ts)
}

func BenchmarkPushFront(t *testing.B) {
	q.PushFront(ts)
}

func BenchmarkFront(t *testing.B) {
	q.Front()
}

func BenchmarkBack(t *testing.B) {
	q.Back()
}

func BenchmarkMutlPushFront(t *testing.B) {
	q.PushFront(ti, ts, ti, ts, ti, ts, ti, ts, ti, ts)
}

func BenchmarkMutlPushBack(t *testing.B) {
	q.PushBack(ti, ts, ti, ts, ti, ts, ti, ts, ti, ts)
}
