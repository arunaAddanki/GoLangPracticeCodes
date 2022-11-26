package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type singlyLinkedList struct {
	head *Node
	tail *Node
	size int
}

func (sl *singlyLinkedList) len() int {
	return sl.size
}

func (sl *singlyLinkedList) isempty() bool {
	return sl.size == 0
}

func (sl *singlyLinkedList) addlast(n int) {
	new := &Node{n, nil}
	if sl.isempty() {
		sl.head = new
		sl.tail = new
	} else {
		sl.tail.next = new
		sl.tail = new
	}
	sl.size++
}

func (sl *singlyLinkedList) display() {
	p := sl.head
	for p != nil {
		fmt.Print(p.data, "-->")
		p = p.next
	}
	fmt.Println()
}

func main() {
	a := singlyLinkedList{}
	a.addlast(10)
	a.addlast(20)
	a.addlast(30)
	a.display()
}
