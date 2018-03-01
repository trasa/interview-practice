package main

import (
	"fmt"
	"github.com/golang-collections/collections/stack"
)

func doreversetest() {
	a := &LinkedListNode{value: "A"}
	b := &LinkedListNode{value: "B"}
	c := &LinkedListNode{value: "C"}
	d := &LinkedListNode{value: "D"}
	e := &LinkedListNode{value: "E"}
	a.next = b
	b.next = c
	c.next = d
	d.next = e

	printList(reverseLinkedListUsingExtraStorage(a))
	printList(reverseLinkedListInPlace(a))

	onenode := &LinkedListNode{value: "one"}
	printList(reverseLinkedListUsingExtraStorage(onenode))
	printList(reverseLinkedListInPlace(onenode))

	emptyList := reverseLinkedListInPlace(nil)
	fmt.Printf("emptylist: %t\n", emptyList == nil)
	fmt.Printf("emptylist: %t\n", reverseLinkedListUsingExtraStorage(nil) == nil)

}

// reverse a linked list in place
func reverseLinkedListInPlace(head *LinkedListNode) *LinkedListNode {
	//  a -> b -> c
	//^prev
	//  ^ cur

	var prev *LinkedListNode

	for current := head; current != nil; {
		// save the old next, which will become previous
		next := current.next
		current.next = prev
		// prepare next loop
		prev = current
		current = next
	}
	return prev

}

func reverseLinkedListUsingExtraStorage(head *LinkedListNode) *LinkedListNode {
	if head == nil {
		return nil
	}

	s := stack.New()
	for current := head; current != nil; current = current.next {
		s.Push(current)
	}

	newHead := s.Pop().(*LinkedListNode)
	for current := newHead; current != nil; current = current.next {
		if s.Len() > 0 {
			current.next = s.Pop().(*LinkedListNode)
		} else {
			current.next = nil
		}
	}
	return newHead
}
