package tree

import (
	"fmt"
	"testing"
)

func TestTreeToList(t *testing.T) {
	root := &BtreeNode[string]{Node: "root"}
	l1 := &BtreeNode[string]{Node: "l1"}
	r1 := &BtreeNode[string]{Node: "r1"}
	root.Left = l1
	root.Right = r1
	n1 := &BtreeNode[string]{Node: "n1"}
	n2 := &BtreeNode[string]{Node: "n2"}
	n3 := &BtreeNode[string]{Node: "n3"}
	l1.Left = n1
	r1.Right = n2
	n1.Left = n3
	treeList, err := BtreeToList(root)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
	fmt.Println(treeList)
	fmt.Println(len(treeList))
}
