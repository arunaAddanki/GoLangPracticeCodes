//method1
package main
import "fmt"

type Node struct {
	data int
	next *Node
}

type CircularLinkedList struct {
	head *Node
	tail *Node
	size int
}

func (cl *CircularLinkedList) len() int {
	return cl.size
}

func (cl *CircularLinkedList) isempty() bool {
	return cl.size == 0
}

func (cl *CircularLinkedList) search(v int) int {
	p := cl.head
	i := 0
	for i < cl.len() {
		if p.data == v {
			return i
		}
		p = p.next
		i++
	}
	return -1
}

func (cl *CircularLinkedList) addlast(n int) {
	new := &Node{n, nil}
	if cl.isempty() {
		new.next = new
		cl.head = new
	} else {
		new.next = cl.head
		// new.next = cl.tail.next
		cl.tail.next = new
	}
	cl.tail = new
	cl.size++
}

func (cl *CircularLinkedList) addany(n int, index int) {
	if index == 0 {
		cl.addfirst(n)
	} else if index > (cl.len())-1 {
		cl.addlast(n)
	} else {
		new := &Node{n, nil}
		p := cl.head
		i := 1
		for i < index {
			p = p.next
			i++
		}
		new.next = p.next
		p.next = new
		cl.size++
	}
}

func (cl *CircularLinkedList) addfirst(n int) {
	new := &Node{n, nil}
	if cl.isempty() {
		new.next = new
		cl.tail = new
	} else {
		new.next = cl.head
	}
	cl.head = new
	cl.size++
}

func (cl *CircularLinkedList) removefirst() {
	if cl.isempty() {
		fmt.Println("Circular Linked LIst is empty")
	} else {
		cl.head = cl.head.next
		cl.tail.next = cl.head
		cl.size--
	}
	if cl.isempty() {
		cl.head = nil
		cl.tail = nil
	}
}

func (cl *CircularLinkedList) removelast() {
	if cl.isempty() {
		fmt.Println("Circular Linked lis is empty")
	} else {
		p := cl.head
		i := 1
		for i < (cl.len() - 1) {
			p = p.next
			i++
		}
		p.next = cl.tail.next
		cl.tail = p
		cl.size--
	}
	if cl.isempty() {
		cl.head = nil
		cl.tail = nil
	}
}

func (cl *CircularLinkedList) removeany(index int) {
	if index > (cl.len())-1 {
		fmt.Println("Index out of range")
	} else if index == (cl.len())-1 {
		cl.removelast()
	} else if index == 0 {
		cl.removefirst()
	} else {
		p := cl.head
		i := 1
		for i < index {
			p = p.next
			i++
		}
		p.next = p.next.next
		cl.size--
	}
	if cl.isempty() {
		cl.head = nil
		cl.tail = nil
	}
}

func (cl *CircularLinkedList) display() {
	p := cl.head
	i := 0
	for i < cl.len() {
		fmt.Print(p.data, "-->")
		p = p.next
		i++
	}
	fmt.Println()
	fmt.Println(cl.tail.next.data)
}

func main() {
	cll := CircularLinkedList{}
	cll.addlast(10)
	cll.addlast(20)
	cll.addlast(30)
	cll.addfirst(40)
	cll.addfirst(50)
	cll.addfirst(60)
	cll.addany(10, 2)
	cll.addlast(25)
	cll.display()
	// fmt.Println(cll.search(100))
	fmt.Println(cll.len())
	// cll.removefirst()
	// cll.removelast()
	// cll.removeany(cll.len())
	// cll.removeany(cll.len()-1)
	cll.removeany(2)
	cll.display()
	fmt.Println(cll.len())

}




//method2

// package main
// import "fmt"

// // import "fmt";import "errors"
// // Node represents a node of linked list

// type CNode struct {
// 	value int
// 	next *CNode
// }

// // LinkedList represents a linked list
// type CirSingLinkedList struct {
// 	head *CNode
// 	tail *CNode
// 	len int
// }

// func (cl *CirSingLinkedList) IsEmpty() bool{
// 	if(cl.len==0){
// 		return true
// 	}else{
// 		return false
// 	}
// }
// func (cl *CirSingLinkedList) Addfirst(v int){
// 	n := CNode{value: v}
// 	if(cl.len==0){
// 		n.next = &n
// 		cl.head = &n
// 		cl.tail = &n
// 	}else{
// 		n.next = cl.head
// 		cl.head = &n
// 		cl.tail.next = &n

// 	}
// 	cl.len++
// 	return
// }
// func (cl *CirSingLinkedList) AddLast(v int){
// 	n := CNode{value: v}
// 	if(cl.len==0){
// 		n.next = &n
// 		cl.head = &n
// 		cl.tail = &n
// 	}else{
// 		n.next = cl.tail.next
// 		cl.tail.next = &n
// 	}
// 	cl.tail = &n
// 	cl.len++
// 	return
// }
// func (cl *CirSingLinkedList) AddAny(pos int, val int){
// 	n := CNode{value: val}
// 	if pos < 0 {
// 		return
// 	}
// 	if pos == 0 {
// 		cl.head = &n
// 		cl.len++
// 		return
// 	}
// 	if pos > cl.len {
// 		return
// 	}
// 	new := cl.GetPos(pos)
// 	n.next = new
// 	prevNode := cl.GetPos(pos - 1)
// 	prevNode.next = &n
// 	cl.len++
// }
// func (cl *CirSingLinkedList) PrintAll(){
// 	if(cl.len==0){
// 		fmt.Println("No nodes there")
// 		return
// 	}
// 	n := cl.head
// 	i:=0	
// 	for(i<cl.len){
// 		fmt.Println(n.value)
// 		n = n.next
// 		i++
// 	}
	
// }
// func (cl *CirSingLinkedList) GetPos(pos int) *CNode {
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
// func (cl *CirSingLinkedList) Size() int{
// 	return cl.len
// }
// func (cl *CirSingLinkedList) RemoveFirst(){
// 	if(cl.IsEmpty()){
// 		fmt.Println("No elements is there to remove....")
// 		return
// 	}
// 	fmt.Println("The removed elem is -", cl.head.value )
// 	cl.tail.next = cl.head.next
// 	cl.head = cl.head.next
// 	if(cl.IsEmpty()){
// 		cl.head = nil
// 		cl.tail = nil
// 	}
// 	cl.len--
// }
// func (cl *CirSingLinkedList) RemoveLast(){
// 	if(cl.IsEmpty()){
// 		fmt.Println("No elements is there to remove....")
// 		return
// 	}
// 	fmt.Println("The removed elem is -", cl.tail.value )
// 	n := cl.head
// 	i:=0	
// 	for(i<cl.len-1){
// 		n = n.next
// 		i++
// 	}
// 	cl.tail = n
// 	cl.tail.next = cl.head

// 	cl.len--
// }
// func (cl *CirSingLinkedList) RemoveAny(pos int){
// 	n := cl.head
// 	i := 1
// 	for(i<pos-1){
// 		n = n.next
// 		i++
// 	}
// 	fmt.Println("The removed elemnt is- ",n.next.value)
// 	n.next = n.next.next
// 	cl.len--
// }

// func main(){
// 	cll := CirSingLinkedList{}
// 	fmt.Println(cll.IsEmpty())
// 	cll.Addfirst(10)
// 	cll.Addfirst(9)
// 	cll.Addfirst(8)
// 	cll.Addfirst(7)
// 	cll.Addfirst(6)
// 	cll.Addfirst(5)
// 	cll.Addfirst(4)
// 	cll.Addfirst(3)
// 	cll.Addfirst(2)
// 	cll.Addfirst(1)
// 	fmt.Println()
// 	fmt.Println(cll.Size())
// 	cll.RemoveFirst()
// 	cll.RemoveLast()
// 	cll.RemoveAny(5)
// 	cll.PrintAll()
// 	fmt.Println(cll.Size())

// }

