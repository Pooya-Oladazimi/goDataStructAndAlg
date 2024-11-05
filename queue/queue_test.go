package queue

import "testing"

func TestEnqueueToEmptyQueue(t *testing.T) {
	q := Newqueue[int](3)
	_, err := q.enqueue(100)
	if err != nil {
		t.Fatalf("%v", err)
	}
}

func TestEnqueueToFullQueue(t *testing.T) {
	q := Newqueue[int](2)
	_, err := q.enqueue(100)
	_, err = q.enqueue(100)
	_, err = q.enqueue(200)
	if err == nil {
		t.Fatalf("full queue %v accepted a new element: %v", q.q, 200)
	}
}

func TestDequeueFromNonEmptyQueue(t *testing.T) {
	q := Newqueue[int](3)
	_, _ = q.enqueue(100)
	_, _ = q.enqueue(200)
	_, _ = q.enqueue(300)
	element, errd := q.dequeue()
	if element != 100 || errd != nil {
		t.Fatalf("dequeue failed for a non empty queue.")
	}
}

func TestDequeueFromEmptyQueue(t *testing.T) {
	q := Newqueue[int](3)
	q.enqueue(100)
	q.enqueue(200)
	q.enqueue(300)
	_, errd := q.dequeue()
	_, errd = q.dequeue()
	_, errd = q.dequeue()
	_, errd = q.dequeue()
	if errd == nil {
		t.Fatalf("dequeue did not fail for an empty queue.")
	}
}
