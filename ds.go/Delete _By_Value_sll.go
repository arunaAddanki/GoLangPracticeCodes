package main

import (
	"fmt"
)

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l linkedList) printListData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d\t", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
}

func (l *linkedList) deleteWithValue(value int) {
	previousToDelete := l.head                          // here we are comparing the data with next node address why because its single linkedlist we only having next node address so
	for previousToDelete.next.data != value {
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next  //once find the node to be deleted we have to skip it and delete it
	l.length--                                          // we are updating length here
}

func main() {
	mylist := linkedList{}
	node1 := &node{data: 1000}
	node2 := &node{data: 2000}
	node3 := &node{data: 3000}
	node4 := &node{data: 4000}
	node5 := &node{data: 5000}
	node6 := &node{data: 6000}
	node7 := &node{data: 7000}
	node8 := &node{data: 8000}
	node9 := &node{data: 9000}
	node10 := &node{data: 10000}

	mylist.prepend(node1)
	mylist.prepend(node2)
	mylist.prepend(node3)
	mylist.prepend(node4)
	mylist.prepend(node5)
	mylist.prepend(node6)
	mylist.prepend(node7)
	mylist.prepend(node8)
	mylist.prepend(node9)
	mylist.prepend(node10)
    


	mylist.printListData()
	mylist.deleteWithValue(9000)
	fmt.Printf("\n")
	mylist.printListData()


}
