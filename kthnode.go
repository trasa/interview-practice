package main

import (
	"github.com/golang-collections/collections/stack"
	"fmt"
)

type LinkedListNode struct {
	value string
	next *LinkedListNode
}


func kthNodeTest() {
	a := &LinkedListNode{ value: "Angel Food"}
	b := &LinkedListNode{ value: "Bundt"}
	c := &LinkedListNode{ value: "Cheese"}
	d := &LinkedListNode{ value: "Devil's Food"}
	e := &LinkedListNode{ value: "Eccles"}
	a.next = b
	b.next = c
	c.next = d
	d.next = e

	last := kthToLastNodeStack(2, a)
	fmt.Printf("%s\n", last.value)

	last = kthToLastNoAlloc(2, a)
	fmt.Printf("%s\n", last.value)

	last = kthToLastStick(2, a)
	fmt.Printf("%s\n", last.value)
}

func kthToLastNoAlloc(k int, node *LinkedListNode) (result *LinkedListNode) {
	// get length
	// don't forget to count the head!
	len := 1
	for n := node; n != nil; n = n.next {
		len++
	}

	// then walk to (len - k) + 1 node
	n := node
	for i := 0; i < len - k; i++ {
		result = n
		n = n.next
	}
	return
}

func kthToLastStick(k int, node *LinkedListNode) (result *LinkedListNode) {
	left := node
	right := node

	// move right to the kth element
	for i := 0; i < k - 1; i++ {
		right = right.next
	}

	// move right and left, keeping k distance, until right = end of list
	for right.next != nil {
		left = left.next
		right = right.next
	}
	return left

}

func kthToLastNodeStack(k int, node *LinkedListNode) (result *LinkedListNode) {
	s := stack.New()

	for n := node; n != nil; n = n.next {
		s.Push(n)
	}

	for i := 0; i < k; i++ {
		result = s.Pop().(*LinkedListNode)
	}
	return
}