package tree

import (
	"Pooya-Oladazimi/goDataStructAndAlg/queue"
	"errors"
)

type BtreeNode[T any] struct {
	Node  T
	Left  *BtreeNode[T]
	Right *BtreeNode[T]
}

func BtreeToList[T any](root *BtreeNode[T]) ([]T, error) {
	res := make([]T, 0)
	if root == nil {
		return res, errors.New("tree is empty")
	}

	var firstEmptyElement T
	res = append(res, firstEmptyElement)
	q := queue.Newqueue[*BtreeNode[T]](-1)
	q.Enqueue(root)
	for !q.IsEmpty() {
		cnode, err := q.Dequeue()
		checkError(err)
		res = append(res, cnode.Node)
		if cnode.Left != nil {
			q.Enqueue(cnode.Left)
		} else {
			var empty T
			res = append(res, empty)
		}
		if cnode.Right != nil {
			q.Enqueue(cnode.Right)
		} else {
			var empty T
			res = append(res, empty)
		}

	}
	return res, nil
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
