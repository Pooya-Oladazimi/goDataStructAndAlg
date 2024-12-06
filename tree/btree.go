package tree

import (
	"Pooya-Oladazimi/goDataStructAndAlg/queue"
	"errors"
	"fmt"
	"math"
	"strings"
)

type BtreeNode[T any] struct {
	Node  T
	Left  *BtreeNode[T]
	Right *BtreeNode[T]
}

func (bt *BtreeNode[T]) Height() (int, error) {
	if bt == nil {
		return -1, errors.New("node is null.")
	}
	treeList, err := bt.ToList()
	treeNodeCount := 0
	if err == nil {
		treeNodeCount = len(treeList) - 1
	}
	height := 0
	for treeNodeCount >= 2 {
		treeNodeCount /= 2
		height += 1
	}
	return height, nil
}

func (bt *BtreeNode[T]) pprint() {
	if bt == nil {
		fmt.Println("either the tree is empty or the node is a leaf.")
		return
	}
	treeHeight, _ := bt.Height()
	level := 0
	q := queue.Newqueue[*BtreeNode[T]](-1)
	q.Enqueue(bt)
	indent := treeHeight * 5
	for level <= treeHeight {
		levelNodeCap := int64(math.Pow(2.0, float64(level)))
		line1 := strings.Repeat(" ", indent-(level*2))
		line2 := ""
		if level != treeHeight {
			line2 = strings.Repeat(" ", indent-(level*2)-4)
		}
		for levelNodeCap > 0 {
			cNode, err := q.Dequeue()
			checkError(err)
			if cNode == nil {
				levelNodeCap -= 1
				q.Enqueue(nil)
				q.Enqueue(nil)
				line1 += "  "
				continue
			}
			line1 += fmt.Sprintf("%v", cNode.Node) + " "
			if cNode.Left != nil {
				line2 += "  / "
				q.Enqueue(cNode.Left)
			} else {
				line2 += "  "
				q.Enqueue(nil)
			}
			if cNode.Right != nil {
				line2 += "\\"
				q.Enqueue(cNode.Right)
			} else {
				line2 += " "
				q.Enqueue(nil)
			}
			levelNodeCap -= 1
		}
		fmt.Println(line1)
		fmt.Println(line2)
		level += 1
	}
}

func (bt *BtreeNode[T]) ToList() ([]T, error) {
	res := make([]T, 0)
	if bt == nil {
		return res, errors.New("tree is empty or node is leaf.")
	}

	var firstEmptyElement T
	res = append(res, firstEmptyElement)
	q := queue.Newqueue[*BtreeNode[T]](-1)
	q.Enqueue(bt)
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
