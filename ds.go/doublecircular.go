//method1
package main
import "fmt"

type Node struct {
	prev  *Node
	value int
	next  *Node
}

type DoubleCircular struct {
	head *Node
	tail *Node
	size int
}

func (dc *DoubleCircular) addLast(val int) {
	newNode := &Node{value: val}
	if dc.size == 0 {
		newNode.next = newNode
		dc.head = newNode
		dc.head.prev = newNode
	} else {
		newNode.next = dc.tail.next
		dc.tail.next = newNode
		newNode.prev = dc.tail
	}
	dc.tail = newNode
	dc.size++
}

func (dc *DoubleCircular) display() {
	ptr := dc.head
	for i := 0; i < dc.size; i++ {
		fmt.Printf("%v--->", ptr)
		ptr = ptr.next
	}
	fmt.Println()
}

func (dc *DoubleCircular) addFirst(val int) {
	newNode := &Node{value: val}
	if dc.size == 0 {
		newNode.next = dc.head
		dc.head = newNode
		dc.head.prev = newNode
	} else {
		dc.tail.next = newNode
		newNode.next = dc.head
		newNode.prev = dc.tail
		dc.head = newNode

	}
	dc.size++
}

func (dc *DoubleCircular) addBetween(val, pos int) {
	newest := &Node{value: val}
	ptr := dc.head
	for i := 1; i < pos-1; i++ {
		ptr = ptr.next
	}
	newest.next = ptr.next
	newest.prev = ptr
	ptr.next = newest

	dc.size++
}

func (dc *DoubleCircular) delFirst() {
	if dc.size == 0 {
		fmt.Println("list is empty.")
	} else {
		dc.tail.next = dc.head.next
		dc.head = dc.head.next
		dc.head.prev = dc.tail
	}
	dc.size--
}

func (dc *DoubleCircular) delLast() {
	ptr := dc.head
	if dc.size == 0 {
		fmt.Println("list is empty.")
	} else {
		for i := 0; i < dc.size; i++ {
			ptr = ptr.next
		}
		dc.head.prev = ptr.next
		ptr.next = dc.head
	}
	dc.size--
}

func main() {
	dc := DoubleCircular{}
	dc.addLast(1)
	dc.display()
	dc.addLast(2)
	dc.display()
	dc.addLast(3)
	dc.display()
	dc.addLast(4)
	dc.display()
	// dc.addFirst(0)
	// dc.display()
	// dc.addBetween(33, 3)
	// dc.display()
}



//method2

// package main
// import "fmt"

// type CDNode struct {
// 	data int
// 	prev *CDNode
// 	next *CDNode
// }

// type CirdoubLinkedList struct {
// 	len  int
// 	tail *CDNode
// 	head *CDNode
// }
// func (d *CirdoubLinkedList) IsEmpty() bool{
// 	if(d.len==0){
// 		return true
// 	}else{
// 		return false
// 	}
// }
// func(d *CirdoubLinkedList) AddFirst(val int){
// 	n := CDNode{data: val}
// 	if d.IsEmpty(){
// 		d.head = &n
// 		d.tail = &n
// 		d.len++
// 		return
// 	}
// 	n.next = d.head
// 	d.head.prev = &n
// 	d.head = &n
// 	d.tail.next = d.head
// 	d.len++
// 	return
// }
// func (d *CirdoubLinkedList) RemoveFirst(){
// 	if d.IsEmpty(){
// 		fmt.Println("No nodes to remove")
// 		return
// 	}
// 	fmt.Println("The removed element is : ", d.head.data)
// 	d.tail.next = d.head.next
// 	d.head = d.head.next
// 	d.len--
// 	return
// }
// func (d *CirdoubLinkedList) AddLast(val int){
// 	n := CDNode{data: val}
// 	if d.IsEmpty(){
// 		d.head = &n
// 		d.tail = &n
// 		d.len++
// 		return
// 	}
// 	n.prev = d.tail
// 	d.tail.next = &n
// 	d.tail = &n
// 	d.tail.next = d.head
// 	d.len++
// 	return
// }
// func (d *CirdoubLinkedList) RemoveLast(){
// 	if d.IsEmpty(){
// 		fmt.Println("No nodes to remove")
// 		return
// 	}
// 	fmt.Println("The removed element is- ",d.tail.data)
// 	d.tail.prev.next = d.head
// 	d.len--
// 	return
// }
// func (d *CirdoubLinkedList) AddAny(pos int,val int){
// 	n := CDNode{data: val}
// 	if d.IsEmpty(){
// 		d.head = &n
// 		d.tail = &n
// 		d.len++
// 		return
// 	}
// 	ptr := d.GetPos(pos)
// 	ptr.prev = &n
// 	n.next = ptr
// 	prev := d.GetPos(pos-1)
// 	n.prev = prev
// 	prev.next = &n
// 	d.len++
// 	return
// }
// func (d *CirdoubLinkedList) RemoveAny(pos int){
// 	if d.IsEmpty(){
// 		fmt.Println("no nodes to renove")
// 		return
// 	}
// 	ptr := d.GetPos(pos)
// 	fmt.Println("The removed element is : ", ptr.data)
// 	ptr.prev.next = ptr.next
// 	d.len--
// 	return
// }
// func (cl *CirdoubLinkedList) GetPos(pos int) *CDNode {
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
// func (d *CirdoubLinkedList) PrintAll(){
// 	if(d.len==0){
// 		fmt.Println("No nodes in list")
// 		return
// 	}
// 	n := d.head
// 	for i:=0; i<d.len;i++{
// 		fmt.Println(n.data)
// 		n = n.next
// 	}
// }
// func (d *CirdoubLinkedList) Reverse(){
// 	if d.IsEmpty(){
// 		fmt.Println("No node is there")
// 		return
// 	}
// 	n := d.tail
// 	for i:=0;i<d.len;i++{
// 		fmt.Println(n.data)
// 		n = n.prev
// 	}
// }
// func (d *CirdoubLinkedList) Len() int{
// 	return d.len
// }
// func (d *CirdoubLinkedList) Search(val int){
// 	if(d.len==0){
// 		fmt.Println("no nodes there to search")
// 	}
// 	n := d.head
// 	for i:=0;i<d.len;i++{
// 		if(n.data==val){
// 			fmt.Println("Serached item found at :", i)
// 		}
// 		n = n.next
// 	}

// } 

// func main(){
// 	dc := CirdoubLinkedList{}
	
// 	dc.AddLast(50)
// 	dc.AddLast(51)
// 	dc.AddLast(52)
// 	dc.AddAny(1,60)
// 	fmt.Println(dc.len)
// 	dc.PrintAll()
// 	dc.Search(60)
// 	fmt.Println("-----------------")
// 	dc.RemoveAny(1)
// 	fmt.Println(dc.len)
// 	dc.PrintAll()
// }