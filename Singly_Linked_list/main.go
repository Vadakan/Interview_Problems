package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head   *Node
	length int
}

func (l *LinkedList) prepend(n *Node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l LinkedList) printListData() {
	toprint := l.head
	for l.length != 0 {
		fmt.Printf("%v ", toprint.data)
		toprint = toprint.next
		l.length--
	}
	fmt.Printf("\n")
}

func (l *LinkedList) deleteWithValue(value int) {
	if l.length == 0 {
		return
	}
	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}
	previoustodelete := l.head
	for previoustodelete.next.data != value {
		if previoustodelete.next.next == nil {
			return
		}
		previoustodelete = previoustodelete.next
	}
	previoustodelete.next = previoustodelete.next.next
	l.length--

}

func main() {

	mylist := LinkedList{}
	node1 := &Node{data: 48}
	node2 := &Node{data: 18}
	node3 := &Node{data: 16}
	node4 := &Node{data: 2}
	node5 := &Node{data: 7}
	node6 := &Node{data: 5}
	node7 := &Node{data: 38}
	mylist.prepend(node1)
	mylist.prepend(node2)
	mylist.prepend(node3)
	mylist.prepend(node4)
	mylist.prepend(node5)
	mylist.prepend(node6)
	mylist.prepend(node7)

	fmt.Println(mylist)
	mylist.printListData()
	mylist.deleteWithValue(48)
	mylist.deleteWithValue(100)
	mylist.deleteWithValue(2)
	emptylist := LinkedList{}
	emptylist.deleteWithValue(10)
	mylist.printListData()

}

