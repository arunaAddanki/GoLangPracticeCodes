//method1

package main
import "fmt"

type Node struct {
	data int
	next *Node
}

type SinglyLinkedList struct {
	head *Node
	tail *Node
	size int
}

func (sl *SinglyLinkedList) len() int {
	return sl.size
}

func (sl *SinglyLinkedList) addlast(n int) {
	new := &Node{n, nil}
	if sl.head == nil {
		sl.head = new
	} else {
		sl.tail.next = new
	}
	sl.tail = new
	sl.size++
}

func (sl *SinglyLinkedList) addfirst(n int) {
	new := &Node{n, nil}
	if sl.head == nil {
		sl.head = new
		sl.tail = new
	} else {
		new.next = sl.head
		sl.head = new
	}
	sl.size++
}

func (sl *SinglyLinkedList) addany(n int, index int) {
	new := &Node{n, nil}
	p := sl.head
	count := 0
	for count < index {
		p = p.next
		count++
	}
	new.next = p.next
	p.next = new
	sl.size++
}

func (sl *SinglyLinkedList) removefirst() {
	if sl.head == nil {
		fmt.Println("Linked list is empty")
	} else {
		p := sl.head.next
		sl.head = p
		sl.size--
	}
	if sl.head == nil {
		sl.tail = nil
	}
}

func (sl *SinglyLinkedList) removelast() {
	if sl.head == nil {
		fmt.Println("Linked list is empty")
	} else {
		p := sl.head
		e := 1
		for e < (sl.len())-1 {
			p = p.next
			e++
		}
		sl.tail = p
		sl.tail.next = nil
		sl.size--
	}
}

func (sl *SinglyLinkedList) removeany(i int) {
	if i != 0 && i != (sl.len())-1 && i < sl.len() {
		if sl.head == nil {
			fmt.Println("Linked list is empty")
		} else {
			p := sl.head
			count := 1
			for count < i {
				p = p.next
				count++
			}
			p.next = p.next.next
			sl.size--
		}
	} else if i == 0 {
		sl.removefirst()
	} else if i == (sl.len())-1 {
		sl.removelast()
	} else {
		fmt.Println("Index ot of range...!")
	}
}

func (sl *SinglyLinkedList) search(key int) int {
	p := sl.head
	index := 0
	for p != nil {
		if p.data == key {
			return index
		}
		p = p.next
		index++
	}
	return -1
}

func (sl *SinglyLinkedList) display() {
	p := sl.head
	for p != nil {
		fmt.Print(p.data, "-->")
		p = p.next
	}
}

func main() {
	l := &SinglyLinkedList{}
	l.addlast(10)
	l.addlast(20)
	l.addlast(30)
	l.display()
	fmt.Println()
	fmt.Println(l.len())
	// l.removefirst()
	// l.addany(40,3)
	l.removeany(0)
	l.display()
	fmt.Println()
	fmt.Println(l.len())

}





//method2


// package main
// import (
// 	"fmt"
// )

// type node struct {
// 	data int
// 	next *node
// }

// // LinkedList is a linked list
// type linkedList struct {
// 	length int
// 	head   *node
// 	tail   *node
// }

// // Len Function returns Length of the LinkedList
// func (l *linkedList) Len() int {
// 	return l.length
// }

// // PushBack Function inserts a new node at the end of the LinkedList
// func (l *linkedList) PushBack(n *node) {
// 	if l.head == nil {
// 		l.head = n
// 		l.tail = n
// 		l.length++
// 	} else {
// 		l.tail.next = n
// 		l.tail = n
// 		l.length++
// 	}
// }

// func (l linkedList) Display() {
// 	for l.head != nil {
// 		fmt.Printf("%v -> ", l.head.data)
// 		l.head = l.head.next
// 	}
// 	fmt.Println()
// }

// func (l linkedList) Front() (int, error) {
// 	if l.head == nil {
// 		return 0, fmt.Errorf("Cannot Find Front Value in an Empty linked list")
// 	}
// 	return l.head.data, nil
// }

// func (l linkedList) Back() (int, error) {
// 	if l.head == nil {
// 		return 0, fmt.Errorf("Cannot Find Front Value in an Empty linked list")
// 	}
// 	return l.tail.data, nil
// }

// func (l *linkedList) Reverse() {
// 	curr := l.head
// 	l.tail = l.head
// 	var prev *node
// 	for curr != nil {
// 		temp := curr.next
// 		curr.next = prev
// 		prev = curr
// 		curr = temp
// 	}
// 	l.head = prev
// }

// func (l *linkedList) Delete(key int) {

// 	if l.head.data == key {
// 		l.head = l.head.next
// 		l.length--
// 		return
// 	}
// 	var prev *node = nil
// 	curr := l.head
// 	for curr != nil && curr.data != key {
// 		prev = curr
// 		curr = curr.next
// 	}
// 	if curr == nil {
// 		fmt.Println("Key Not found")
// 		return
// 	}
// 	prev.next = curr.next
// 	l.length--
// 	fmt.Println("Node Deleted")

// }

// func main() {
// 	list := linkedList{}
// 	node1 := &node{data: 20}
// 	node2 := &node{data: 30}
// 	node3 := &node{data: 40}
// 	node4 := &node{data: 50}
// 	node5 := &node{data: 70}
// 	list.PushBack(node1)
// 	list.PushBack(node2)
// 	list.PushBack(node3)
// 	list.PushBack(node4)
// 	list.PushBack(node5)
// 	fmt.Println("Length = ", list.Len())
// 	list.Display()
// 	list.Delete(40)
// 	list.Reverse()
// 	fmt.Println("Length = ", list.Len()) 
// 	list.Display()
// 	front, _ := list.Front()
// 	back, _ := list.Back()
// 	fmt.Println("Front = ", front)
// 	fmt.Println("Back = ", back)
// }




//method3


// package main
// import "fmt";import "errors"

// // Node represents a node of linked list
// type Node struct {
// 	value int
// 	next *Node
// }

// // LinkedList represents a linked list
// type LinkedList struct {
// 	head *Node
// 	len int
// }
// func (l * LinkedList) Reverse() {
// 	if l.len==0{
// 		fmt.Println("no elements in linked list")
// 		return
// 	}
// 	curr := l.head
// 	var Prev *Node
// 	var Next *Node
// 	Prev = nil
// 	Next = nil
// 	for curr!=nil{
// 		Next = curr.next
// 		curr.next = Prev
// 		Prev = curr
// 		curr = Next
// 	}
// 	l.head = curr
// }

// func (l *LinkedList) Print() {
// 	if l.len == 0 {
// 		fmt.Println("No nodes in list")
// 	}
// 	ptr := l.head
// 	for i := 0; i < l.len; i++ {
// 		fmt.Printf("%d ", ptr.value)
// 		ptr = ptr.next
// 	}
// 	println()
// }
// func (l *LinkedList) AddFirst(val int){
// 	n := Node{}
// 	n.value = val
// 	if l.len==0{
// 		l.head=&n
// 		l.len++
// 		return
// 	}
// 	ptr := l.head
// 	n.next = ptr
// 	l.head = &n
// 	l.len++
// }
// func (l *LinkedList) Insert(val int) {
// 	n := Node{value: val}
// 	if l.len == 0 {
// 		l.head = &n
// 		l.len++
// 		return
// 	}
// 	ptr := l.head
// 	for i := 0; i < l.len; i++ {
// 		if ptr.next == nil {
// 			ptr.next = &n
// 			l.len++
// 			return
// 		}
// 		ptr = ptr.next
// 	}
// }

// func (l *LinkedList) Search(val int) int {
// 	ptr := l.head
// 	for i := 0; i < l.len; i++ {
// 		if ptr.value == val {
// 			return i
// 		}
// 		ptr = ptr.next
// 	}
// 	return -1
// }
// func (l *LinkedList) Sort(){
// 	if l.len==0{
// 		fmt.Println("There is No items in linked list...!")
// 		return
// 	}
// 	ptr := l.head
// 	for ptr.next!=nil{
// 		temp := ptr.next
// 		for temp.next!=nil{
// 			if ptr.value>temp.value{
// 				t := temp.value
// 				temp.value = ptr.value
// 				ptr.value = t
// 			}
// 			temp = temp.next
// 		}
// 		if ptr.value>temp.value{
// 			t1 := temp.value
// 			temp.value = ptr.value
// 			ptr.value = t1
// 		}
// 		ptr = ptr.next
// 	}
// }
// func (l *LinkedList) GetAt(pos int) *Node {
// 	ptr := l.head
// 	if pos < 0 {
// 		return ptr
// 	}
// 	if pos > (l.len - 1) {
// 		return nil
// 	}
// 	for i := 0; i < pos; i++ {
// 		ptr = ptr.next
// 	}
// 	return ptr
// }
// func (l *LinkedList) InsertAt(pos int, val int) {
// 	// create a new node
// 	newNode := Node{}
// 	newNode.value = val
// 	// validate the position
// 	if pos < 0 {
// 		return
// 	}
// 	if l.len == 0 {
// 		l.head = &newNode
// 		l.len++
// 		return
// 	}
// 	if pos > l.len {
// 		return
// 	}
// 	curr := l.GetAt(pos)
// 	newNode.next = curr
// 	prevNode := l.GetAt(pos - 1)
// 	prevNode.next = &newNode
// 	l.len++
// }
// func (l *LinkedList) DeleteAt(pos int) error {
// 	// validate the position
// 	if pos < 0 {
// 		fmt.Println("position can not be negative")
// 		return errors.New("position can not be negative")
// 	}
// 	if l.len == 0 {
// 		fmt.Println("No nodes in list")
// 		return errors.New("No nodes in list")
// 	}
// 	prevNode := l.GetAt(pos - 1)
// 	if prevNode == nil {
// 		fmt.Println("Node not found")
// 		return errors.New("Node not found")
// 	}
// 	prevNode.next = l.GetAt(pos+1)
// 	l.len--
// 	return nil
// }
// func (l *LinkedList) DeleteVal(val int) error {
// 	ptr := l.head
// 	if l.len == 0 {
// 		fmt.Println("List is empty")
// 		return errors.New("List is empty")
// 	}
// 	for i := 0; i < l.len; i++ {
// 		if ptr.value == val {
// 			if i > 0 {
// 				prevNode := l.GetAt(i - 1)
// 				prevNode.next = l.GetAt(i).next
// 			} else {
// 				l.head = ptr.next
// 			}
// 			l.len--
// 			return nil
// 		}
// 		ptr = ptr.next
// 	}
// 	fmt.Println("Node not found")
// 	return errors.New("Node not found")
// }
// func (l *LinkedList) size()int{
// 	return l.len
// }
// func main(){
// 	l := LinkedList{}
// 	var choice, pos, val int
// 	var exit string
// loop:
// 	fmt.Println("Choose the action below..!")
// 	fmt.Println("1.Insert 2.Insert_At 3.Delete 4.Display ")
// 	fmt.Scanln(&choice)
// 	switch choice{
// 	case 1:
// 		fmt.Println("Enter the value : ")
// 		fmt.Scanln(&val)
// 		l.Insert(val)
// 	case 2:
// 		fmt.Println("Enter the value : ")
// 		fmt.Scanln(&val)
// 		fmt.Println("Enter the position you want to insert : ")
// 		fmt.Scan(&pos)
// 		l.InsertAt(pos,val)
// 	case 3:
// 		fmt.Println("Enter the position you want to Delete : ")
// 		fmt.Scan(&pos)
// 		l.DeleteAt(pos)
// 	case 4:
// 		l.Print()
// 	default:
// 		fmt.Println("Entered Wrong choice..!")
// 	}
// 	fmt.Println("Do you want to continue(y/n)..")
// 	fmt.Scanln(&exit)
// 	if exit=="y"{
// 		goto loop
// 	}else{
// 		fmt.Println("bye...8")
// 	}
// }