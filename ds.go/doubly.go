//method1
package main
import "fmt"

type Node struct {
	data int
	next *Node
	prev *Node
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
	size int
}

func (dl *DoublyLinkedList) len() int {
	return dl.size
}

func (dl *DoublyLinkedList) isempty() bool {
	return dl.size == 0
}

func (dl *DoublyLinkedList) addfirst(n int) {
	new := &Node{n, nil, nil}
	if dl.isempty() {
		// dl.head = new
		dl.tail = new
	} else {
		dl.head.prev = new
		new.next = dl.head
		// dl.head = new
	}
	dl.head = new
	dl.size++
}

func (dl *DoublyLinkedList) addlast(n int) {
	new := &Node{n, nil, nil}
	if dl.isempty() {
		dl.head = new
		// dl.tail = new
	} else {
		dl.tail.next = new
		new.prev = dl.tail
		// dl.tail = new
	}
	dl.tail = new
	dl.size++
}

func (dl *DoublyLinkedList) addany(n int, index int) {
	if index == 0 {
		dl.addfirst(n)
	} else if index >= (dl.len()) {
		dl.addlast(n)
	} else {
		new := &Node{n, nil, nil}
		p := dl.head
		i := 1
		for i < index {
			p = p.next
			i++
		}
		new.prev = p
		new.next = p.next
		p.next.prev = new
		p.next = new
	}
	dl.size++
}

func (dl *DoublyLinkedList) removefirst() {
	if dl.isempty() {
		fmt.Println("Doubly Linked List is empty")
	} else {
		dl.head = dl.head.next
		dl.size--
	}
	if dl.isempty() {
		dl.head = nil
		dl.tail = nil
	}
}

func (dl *DoublyLinkedList) removelast() {
	if dl.isempty() {
		fmt.Println("Doubly linked list is empty")
	} else {
		dl.tail = dl.tail.prev
		dl.tail.next = nil
		dl.size--
	}
	if dl.isempty() {
		dl.head = nil
		dl.tail = nil
	}
}

func (dl *DoublyLinkedList) removeany(index int) {
	if index >= dl.len() {
		fmt.Println("Index out of range")
	} else if index == 0 {
		dl.removefirst()
	} else if index == (dl.len())-1 {
		dl.removelast()
	} else {
		p := dl.head
		i := 1
		for i < index {
			p = p.next
			i++
		}
		p.next = p.next.next
		p.next.prev = p
		dl.size--
	}
	if dl.isempty() {
		dl.head = nil
		dl.tail = nil
	}
}

func (dl *DoublyLinkedList) display(n int) {
	if n == 0 {
		p := dl.head
		for p != nil {
			fmt.Print(p.data, "-->")
			p = p.next
		}
		fmt.Println()
	} else if n == 1 {
		p := dl.tail
		for p != nil {
			fmt.Print(p.data, "-->")
			p = p.prev
		}
		fmt.Println()
	} else {
		fmt.Println("Invalid argument: arugument must be 0 for forword display or 1 for reverse display")
	}
}

func main() {
	dll := DoublyLinkedList{}
	// 	dll.addfirst(10)
	// 	dll.addfirst(20)
	// 	dll.addfirst(30)
	// 	// dll.display(2)
	// 	dll.addlast(40)
	// 	dll.addlast(50)
	// 	dll.addlast(60)
	// 	dll.addany(11, 2)
	// 	// dll.display(0) // 0 --> forword diplay, 1--> reverse display
	// 	dll.removefirst()
	// 	dll.removefirst()
	// 	dll.display(0)
	// 	// dll.removelast()
	// 	// dll.removelast()
	// 	// dll.display(0)
	// 	// dll.removeany(dll.len() - 1)
	// 	dll.removeany(3)
	// 	dll.display(0)
	// 	fmt.Println(dll.len())
	dll.addfirst(10)
	dll.addfirst(20)
	dll.display(0)
	dll.removeany(0)
	dll.display(0)
}




//method2


// package main
// import "fmt"

// type node struct {
// 	data string
// 	prev *node
// 	next *node
// }

// type doublyLinkedList struct {
// 	len  int
// 	tail *node
// 	head *node
// }

// func initDoublyList() *doublyLinkedList {
// 	return &doublyLinkedList{}
// }

// func (d *doublyLinkedList) AddFirst(data string) {
// 	newNode := &node{
// 		data: data,
// 	}
// 	if d.head == nil {
// 		d.head = newNode
// 		d.tail = newNode
// 	} else {
// 		newNode.next = d.head
// 		d.head.prev = newNode
// 		d.head = newNode
// 	}
// 	d.len++
// 	return
// }

// func (d *doublyLinkedList) AddLast(data string) {
// 	newNode := &node{
// 		data: data,
// 	}
// 	if d.head == nil {
// 		d.head = newNode
// 		d.tail = newNode
// 	} else {
// 		currentNode := d.head
// 		for currentNode.next != nil {
// 			currentNode = currentNode.next
// 		}
// 		newNode.prev = currentNode
// 		currentNode.next = newNode
// 		d.tail = newNode
// 	}
// 	d.len++
// 	return
// }
// func (d *doublyLinkedList) AddAny(pos int, val string){
// 	n := node{data: val}
// 	if d.len==0{
// 		d.head = &n
// 		d.tail = &n
// 		d.len++
// 		return
// 	}
// 	if pos<0{
// 		println("Position cant be negative..!")
// 		return
// 	}
// 	if pos>d.len-1{
// 		println("Position cant be greater than length of the list...!")
// 		return
// 	}
// 	ptr := d.GetPos(pos)
// 	ptr.prev = &n
// 	n.next = ptr
// 	prev := d.GetPos(pos-1)
// 	prev.next = &n
// 	n.prev = prev
// 	d.len++
// 	return
// }
// func (cl *doublyLinkedList) GetPos(pos int) *node {
// 	ptr := cl.head
// 	if pos < 0 {
// 		return ptr
// 	}
// 	if pos > (cl.len - 1) {
// 		return nil
// 	}
// 	for i := 0; i < pos; i++ {
// 		ptr = ptr.next
// 	}
// 	return ptr
// }
// func (d *doublyLinkedList) RemoveFirst(){
// 	if d.len==0{
// 		fmt.Println("List is empty..!")
// 		return
// 	}
// 	fmt.Println("The removed element is : ", d.head.data)
// 	if d.head.next!=nil{
// 		d.head = d.head.next
// 		d.head.prev = nil
// 	}else{
// 		d.head = nil
// 		d.tail = nil
// 	}
// 	d.len--
// }
// func (d *doublyLinkedList) RemoveLast(){
// 	if d.len==0{
// 		fmt.Println("List is empty..!")
// 		return
// 	}
// 	fmt.Println("The removed element is : ", d.tail.data)
// 	if d.tail.prev!=nil{
// 		d.tail = d.tail.prev
// 		d.tail.next = nil
// 	}else{
// 		d.head = nil
// 		d.tail = nil
// 	}
// 	d.len--
// }
// func (d *doublyLinkedList) RemoveAny(pos int){
// 	if d.len==0{
// 		fmt.Println("List is empty..!")
// 		return
// 	}
// 	if pos<0{
// 		println("Position cant be negative..!")
// 		return
// 	}
// 	if pos>d.len-1{
// 		println("Position cant be greater than length of the list...!")
// 		return
// 	}
// 	if d.len==1{
// 		d.head=nil
// 		d.tail = nil
// 		d.len--
// 		return
// 	}
// 	ptr := d.GetPos(pos-1)
// 	fmt.Println("the removed element is : ", ptr.next.data)
// 	next := d.GetPos(pos+1)
// 	ptr.next = ptr.next.next
// 	next.prev = ptr
// 	d.len--
// }
// func (d *doublyLinkedList) TraverseForward() error {
// 	if d.head == nil {
// 		return fmt.Errorf("TraverseError: List is empty")
// 	}
// 	temp := d.head
// 	for temp != nil {
// 		fmt.Printf("%v  ", temp.data)
// 		temp = temp.next
// 	}
// 	fmt.Println()
// 	return nil
// }

// func (d *doublyLinkedList) Reverse() error {
// 	if d.head == nil {
// 		return fmt.Errorf("TraverseError: List is empty")
// 	}
// 	temp := d.tail
// 	for temp != nil {
// 		fmt.Printf("value = %v, prev = %v, next = %v\n", temp.data, temp.prev, temp.next)
// 		temp = temp.prev
// 	}
// 	fmt.Println()
// 	return nil
// }

// func (d *doublyLinkedList) Size() int {
// 	return d.len
// }
// var i string
// func init(){
// 	fmt.Println("Doubly linked list...!")
// 	i = "santhosh"
// }

// func main() {
// 	doublyList := CirSingLinkedList{}
// 	doublyList.AddLast(0)
// 	doublyList.AddLast(0)
// 	doublyList.RemoveLast()
// 	doublyList.PrintAll()
// }