package main                           //prepend means we are going to insert nodes at the frontside of Single Linked List(S.L.L)     
import "fmt"

type node struct{                     //declareing structure of node
	data int                          //node consists of actual value i,e data      
	next *node                        //node consists of address of next node
}

type linkedList struct{               //declaring structure of linkedList
	head *node                        //we have to declare head and length of linkedlist no need to define other things
	length int                        
}

func (l *linkedList) prepend (n *node){     //we want a method to do make operations on S.L.L  (l *linkedlist) is a pointer receiver method to allows us make operations here linkedlist is the receiver method we call prepend to  insert or add nodes at the frontside of S.L.L
     second :=l.head                        //we are shifting data from first node to temporary node second 
	 l.head = n                             //we are creating new node n so head is present in new node this time
	 l.head.next = second                   //here we are pointing the head address to second node we make second node as permenant
	 l.length++                             // we are updating length here
}

func main(){
	mylist :=linkedList{}                  // Here i am assigning linkedlist to mylist
	node1:=&node{data:100}                 //assigning data to nodes.............
	node2:=&node{data:200}
	node3:=&node{data:300}
	node4:=&node{data:400}
	node5:=&node{data:500}

	mylist.prepend(node1)               // i am inserting the actual data here above is declaration of prepend method.......
	mylist.prepend(node2)
	mylist.prepend(node3)
	mylist.prepend(node4)
	mylist.prepend(node5)    
	fmt.Println(mylist)           // printing mylist it will give node5 address and length of list
}