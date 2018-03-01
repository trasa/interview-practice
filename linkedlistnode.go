package main

import "fmt"

type LinkedListNode struct {
	value string
	next  *LinkedListNode
}

func printList(n *LinkedListNode) {

	for ; n != nil; n = n.next {
		fmt.Printf("%s", n.value)
		if n.next != nil {
			fmt.Printf(" -> ")
		}
	}
	fmt.Println()
}
