package tree

import (
	"testing"
)

func getBtreeExample() *BtreeNode[int] {
	/**
	    *           1
	    *          /\
	    *         2  3
	  *          /    \
	  *         4      5
	  *        /
	  *       6
	*/
	root := &BtreeNode[int]{Node: 1}
	l1 := &BtreeNode[int]{Node: 2}
	r1 := &BtreeNode[int]{Node: 3}
	root.Left = l1
	root.Right = r1
	n1 := &BtreeNode[int]{Node: 4}
	n2 := &BtreeNode[int]{Node: 5}
	nX := &BtreeNode[int]{Node: 10}
	n3 := &BtreeNode[int]{Node: 6}
	l1.Left = n1
	r1.Right = n2
	r1.Left = nX
	n1.Left = n3
	return root
}

// func TestTreeToList(t *testing.T) {
// 	root := getBtreeExample()
// 	treeList, err := root.ToList()
// 	if err != nil {
// 		t.Fatalf("Error: %v", err)
// 	}
// 	fmt.Println(treeList)
// 	fmt.Println(len(treeList))
// }

func TestPprint(t *testing.T) {
	root := getBtreeExample()
	root.pprint()
}
