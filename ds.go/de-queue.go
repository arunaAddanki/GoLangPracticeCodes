//BY USING SLICING METHOD

package main
import "fmt"

type node struct {
	data int
	next *node
}
type DEqueue struct {
	len  int
	head,tail *node
	
}

func (q *DEqueue) IsEmpty() bool {
	if q.len == 0 {
		return true
	} else {
		return false
	}
}
func (q *DEqueue) addfirst(e int) {
	new := &node{e, nil}
	if q.head == nil {
		q.head = new
		q.tail = new
		q.len++
	} else {
		new.next = q.head
		q.head = new
		q.len++
	}
}
func (q *DEqueue) addlast(e int) {
	new := &node{e, nil}
	if q.head == nil {
		q.head = new
		q.tail = new
	} else {
		q.tail.next = new
		q.len++
	}
}

func (q *DEqueue) removefirst() {
	if q.head == nil {
		fmt.Println("DEqueuelist is empty")
		return
	}
	q.head = q.head.next
	if q.head == nil {
		q.tail = nil
		q.len--
	}

}
func (q *DEqueue) removelast() {
	if q.head == nil {
		fmt.Println("DEqueuelist is empty")
		return
	}
	q.tail.next = nil
	q.len--

}
func (q *DEqueue) first() {
	if q.IsEmpty(){
		fmt.Println("empty")

	}else{
		fmt.Println(q.head)
	}
	
}
func (q *DEqueue) last() {
	fmt.Println(q.tail)
}

func (q *DEqueue) display() {
	p := q.head
	for p != nil {
		fmt.Print(p.data, "-->")
		p = p.next

	}
}
func main() {
	dque := DEqueue{}
	dque.addlast(30)
	dque.display()
	fmt.Println()
	dque.addfirst(20)
	dque.display()
	fmt.Println()
	dque.addfirst(10)
	dque.display()
	fmt.Println()

	dque.removefirst()
	dque.display()
	fmt.Println()
	dque.addlast(40)
	dque.display()
	fmt.Println()
	dque.removelast()
	dque.display()
	fmt.Println("\n--------------------------------")
	dque.first()
	fmt.Println("--------------------------------")
	dque.last()
}




//BY USING LINKED-LIST METHOD

// package main
// import "fmt"

// type Node struct {
// 	data int
// 	next *Node
// }

// type DeQueue struct {
// 	front *Node
// 	rare  *Node
// 	size  int
// }

// func (dq *DeQueue) len() int {
// 	return dq.size
// }

// func (dq *DeQueue) isempty() bool {
// 	return dq.size == 0
// }

// func (dq *DeQueue) addfirst(n int) {
// 	new := &Node{n, nil}
// 	if dq.isempty() {
// 		dq.front = new
// 		dq.rare = new
// 	} else {
// 		new.next = dq.front
// 		dq.front = new
// 	}
// 	dq.size++
// }

// func (dq *DeQueue) addlast(n int) {
// 	new := &Node{n, nil}
// 	if dq.isempty() {
// 		dq.front = new
// 	} else {
// 		dq.rare.next = new
// 	}
// 	dq.rare = new
// 	dq.size++
// }

// func (dq *DeQueue) removefirst() int {
// 	var r int
// 	if dq.isempty() {
// 		fmt.Println("DEQueue is empty!")
// 	} else {
// 		r = dq.front.data
// 		dq.front = dq.front.next
// 		dq.size--
// 	}
// 	return r
// }

// func (dq *DeQueue) removelast() int {
// 	var r int
// 	if dq.isempty() {
// 		fmt.Println("DEQueue is empty!")
// 	} else {
// 		p := dq.front
// 		i := 1
// 		for i < dq.len()-1 {
// 			p = p.next
// 		}
// 		r = p.next.data
// 		p.next = nil
// 		dq.rare = p
// 		dq.size--
// 	}
// 	return r
// }

// func (dq *DeQueue) first() int {
// 	var r int
// 	if dq.isempty() {
// 		fmt.Println("DEQueue is empty!")
// 	} else {
// 		r = dq.front.data
// 	}
// 	return r
// }

// func (dq *DeQueue) last() int {
// 	var r int
// 	if dq.isempty() {
// 		fmt.Println("DEQueue is empty!")
// 	} else {
// 		r = dq.rare.data
// 	}
// 	return r
// }

// func (dq *DeQueue) display() {
// 	p := dq.front
// 	for p != nil {
// 		fmt.Print(p.data, "<-->")
// 		p = p.next
// 	}
// 	fmt.Println()
// }

// func main() {
// 	dqu := DeQueue{}
// 	dqu.addfirst(10)
// 	dqu.addfirst(20)
// 	dqu.display()
// 	fmt.Println(dqu.len())
// 	dqu.addlast(30)
// 	dqu.addlast(40)
// 	dqu.display()
// 	fmt.Println(dqu.len())
// 	fmt.Println(dqu.removefirst())
// 	dqu.display()
// 	fmt.Println(dqu.removefirst())
// 	dqu.display()
// 	fmt.Println(dqu.removelast())
// 	dqu.display()
// 	dqu.addfirst(10)
// 	dqu.addlast(20)
// 	dqu.display()
// 	fmt.Println(dqu.first())
// 	fmt.Println(dqu.last())
// 	dqu.addfirst(100)
// 	dqu.addlast(300)
// 	fmt.Println(dqu.first())
// 	fmt.Println(dqu.last())
// }

