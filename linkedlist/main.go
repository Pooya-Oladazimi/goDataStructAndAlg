package main

import "fmt"

func main() {
	LL := &LinkedList{head: nil, tail: nil}
	LL.pprint()
	LL.add(100, 0)
	LL.add(200, -1)
	LL.add(300, -1)
	LL.add(400, 1)
	LL.add(420, 2)
	LL.add(430, 1)
	LL.add(440, 1)
	LL.add(500, 1)
	LL.pprint()
	LL.delete(-1)
	LL.delete(0)
	LL.delete(3)
	LL.pprint()
}

type LinkedListNode struct {
	value int
	next  *LinkedListNode
}

type LinkedList struct {
	head *LinkedListNode
	tail *LinkedListNode
}

func (LL *LinkedList) add(val int, position int) {
	if LL.head == nil && (position == -1 || position == 0) {
		Llnode := &LinkedListNode{value: val, next: nil}
		LL.head = Llnode
		LL.tail = Llnode
		return
	} else if LL.head == nil {
		panic("Invalid position to add: List is empty")
	}

	if position == -1 {
		// position -1 means append end of the list
		node := &LinkedListNode{value: val, next: nil}
		LL.tail.next = node
		LL.tail = node
		return
	}

	if position == 0 {
		// head should change
		node := &LinkedListNode{value: val, next: nil}
		node.next = LL.head
		LL.head = node
		return
	}

	if position > 0 {
		cPos := 0
		cNode := LL.head
		for cPos != position {
			if cNode == nil {
				panic("position is out of range.")
			}
			cNode = cNode.next
			cPos++
		}
		node := &LinkedListNode{value: val, next: cNode.next}
		cNode.next = node
		if node.next == nil {
			LL.tail = node
		}
		return
	}

	panic("Invalid position is given")
}

func (LL *LinkedList) delete(position int) {
	if LL.head == nil {
		panic("List is empty.")
	}

	if position == -1 {
		// delete the last element
		if LL.head == LL.tail {
			LL.head = nil
			LL.tail = nil
			return
		}
		cNode := LL.head
		var prevNode *LinkedListNode
		for cNode.next != nil {
			prevNode = cNode
			cNode = cNode.next
		}
		prevNode.next = nil
		LL.tail = prevNode
		return
	}

	if position == 0 {
		if LL.head == LL.tail {
			LL.head = nil
			LL.tail = nil
			return
		}
		LL.head = LL.head.next
		return
	}

	if position > 0 {
		cPos := 0
		cNode := LL.head
		var prevNode *LinkedListNode
		for cPos != position {
			if cNode == nil {
				panic("position is out of range")
			}
			prevNode = cNode
			cNode = cNode.next
			cPos++
		}
		if cNode.next == nil {
			LL.tail = prevNode
		}
		prevNode.next = cNode.next
		return
	}

	panic("Invalid position is given")
}

func (LL *LinkedList) pprint() {
	if LL.head == nil {
		fmt.Println("List is empty")
		return
	}
	cNode := LL.head
	for cNode.next != nil {
		fmt.Printf("%v -> ", cNode.value)
		cNode = cNode.next
	}
	fmt.Printf("%v", cNode.value)
	fmt.Println()
}
