package queue

import "testing"

func TestEnqueueToEmptyQueue(t *testing.T) {
	q := Newqueue[int](3)
	_, err := q.Enqueue(100)
	if err != nil {
		t.Fatalf("%v", err)
	}
}

func TestEnqueueToFullQueue(t *testing.T) {
	q := Newqueue[int](2)
	_, err := q.Enqueue(100)
	_, err = q.Enqueue(100)
	_, err = q.Enqueue(200)
	if err == nil {
		t.Fatalf("full queue %v accepted a new element: %v", q.q, 200)
	}
}

func TestDequeueFromNonEmptyQueue(t *testing.T) {
	q := Newqueue[int](3)
	_, _ = q.Enqueue(100)
	_, _ = q.Enqueue(200)
	_, _ = q.Enqueue(300)
	element, errd := q.Dequeue()
	if element != 100 || errd != nil {
		t.Fatalf("Dequeue failed for a non empty queue.")
	}
}

func TestDequeueFromEmptyQueue(t *testing.T) {
	q := Newqueue[int](3)
	q.Enqueue(100)
	q.Enqueue(200)
	q.Enqueue(300)
	_, errd := q.Dequeue()
	_, errd = q.Dequeue()
	_, errd = q.Dequeue()
	_, errd = q.Dequeue()
	if errd == nil {
		t.Fatalf("Dequeue did not fail for an empty queue.")
	}
}
