
```
package main

import "fmt"

type List struct {
	head *Node
}
type Node struct {
	next *Node
	data int
}

func (l *List) Add(data int) {
	newNode := &Node{data: data}
	if l.head == nil {
		l.head = newNode
	} else {
		current := l.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
}

func (l *List) Push(data int) {
	newHead := &Node{data: data}
	if l.head == nil {
		l.head = newHead
	} else {
		oldHead := l.head
		l.head = newHead
		newHead.next = oldHead
	}
}

func (l *List) Shift() {
	if l.head != nil {
		l.head = l.head.next
	}
}

func (l *List) Pop() {
	if l.head != nil {
		l.head = l.head.next
	}
}

func (l *List) Len() {
	length := 0
	current := l.head
	for current != nil {
		length++
		current = current.next
	}
	fmt.Println("Len =", length)
}

//	func (l *List) DelN(n int) {
//		dummyNode := &Node{next: l.head}
//
//		length := 0
//		curr := dummyNode
//
//		for curr != nil {
//			length++
//			curr = curr.next
//		}
//
//		curr = dummyNode
//		for i := 1; i < n; i++ {
//			curr = curr.next
//		}
//		curr.next = curr.next.next
//	}
func (l *List) DelN(n int) {
	dummyNode := &Node{next: l.head, data: 22}
	curr := dummyNode

	for i := 1; i < n; i++ {
		curr = curr.next
	}
	curr = curr.next.next
}

func main() {
	list := List{}
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.DelN(3)

	current := list.head
	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
}

```
