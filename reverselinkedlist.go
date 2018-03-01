package main

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

	printList(reverseLinkedList(a))

	onenode := &LinkedListNode{value: "one"}
	printList(reverseLinkedList(onenode))
}

// reverse a linked list in place
func reverseLinkedList(head *LinkedListNode) *LinkedListNode {
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
