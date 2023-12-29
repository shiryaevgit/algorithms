package main

import (
	"fmt"
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	list1 := makeList([]int{1, 1, 2, 2, 3, 4})
	res := deleteDuplicates(list1)
	res.Print()
}

func partition(head *ListNode, x int) *ListNode {
	before := &ListNode{}
	after := &ListNode{}
	before_curr := before
	after_curr := after

	for head != nil {
		if head.Val < x {
			before_curr.Next = head
			before_curr = before_curr.Next
		} else {
			after_curr.Next = head
			after_curr = after_curr.Next
		}
		head = head.Next
	}

	after_curr.Next = nil
	before_curr.Next = after.Next

	return before.Next
}

// Remove Duplicates from Sorted List ---------------
func deleteDuplicates(head *ListNode) *ListNode {

	dummyNode := &ListNode{Next: head}
	pred := dummyNode

	for head != nil {
		if head.Next != nil && head.Val == head.Next.Val {
			for head.Next != nil && head.Val == head.Next.Val {
				head = head.Next
			}
			pred.Next = head.Next
		} else {
			pred = pred.Next
		}
		head = head.Next
	}

	return dummyNode.Next
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var res *ListNode = nil

	p1 := headA
	p2 := headB

	for p1.Next != nil && p1 != nil {

		for p2.Next != nil && p2 != nil {
			if p1.Next == p2.Next {
				res = p1
				return res
			} else {
				p2 = p2.Next
			}
		}
		p1 = p1.Next
	}
	return res
}

// mergeKLists -----------------------------
func mergeKLists(lists []*ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy

	for {
		minNode, minIndex := getMinNode(lists)
		if minNode == nil {
			break
		}
		cur.Next = minNode
		lists[minIndex] = lists[minIndex].Next
		cur = cur.Next
	}

	return dummy.Next
}
func getMinNode(lists []*ListNode) (*ListNode, int) {
	min := math.MaxInt32
	var node *ListNode
	var index int

	for i, n := range lists {
		if n == nil {
			continue
		}
		if n.Val < min {
			node = n
			index = i
			min = n.Val
		}
	}
	return node, index
}

// mergeTwoLists -----------------------------
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummyNode := &ListNode{}
	cur := dummyNode

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next
	}

	// Добавляем оставшиеся узлы из одного из списков
	if list1 != nil {
		cur.Next = list1
	} else {
		cur.Next = list2
	}
	return dummyNode.Next
}

// reorderList -----------------------------
func reorderList(head *ListNode) {
	tmp := preMid(head)
	p2 := reversReorder(tmp.Next)

	tmp.Next = nil

	newHead, p1, p1Next := head, head, head

	for p2 != nil {
		p1Next = p1.Next
		p1.Next = p2
		p1 = p2
		p2 = p1Next
	}
	head = newHead
}
func preMid(head *ListNode) *ListNode {
	slow, fast := head, head

	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
func reversReorder(head *ListNode) *ListNode {
	var prev *ListNode = nil
	tmp, cur := head, head

	for cur != nil {
		tmp = cur
		cur = cur.Next
		tmp.Next = prev
		prev = tmp
	}
	return prev
}

// isPalindrome -----------------------------
func isPalindrome(head *ListNode) bool {
	p1 := head
	p2 := mid(head)
	p2 = revrs(p2)

	for p2 != nil {
		if p1.Val != p2.Val {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	return true
}
func mid(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
func revrs(head *ListNode) *ListNode {
	var rev *ListNode = nil
	cur, tmp := head, head

	for cur != nil {
		tmp = cur
		cur = cur.Next
		tmp.Next = rev
		rev = tmp
	}
	return tmp
}

// reverseList -----------------------------
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil

	cur := head
	tmp := cur

	for cur != nil {
		tmp = cur
		cur = cur.Next
		tmp.Next = prev
		prev = tmp
	}
	return prev
}

// removeNthFromEnd -----------------------------
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyNode := &ListNode{Next: head}
	length := 0
	curr := dummyNode

	for curr != nil {
		length++
		curr = curr.Next
	}

	curr = dummyNode
	for i := 0; i < length-n-1; i++ {
		curr = curr.Next
	}
	curr.Next = curr.Next.Next

	return dummyNode.Next
}
func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	dummyNode := &ListNode{Next: head}
	fast := dummyNode

	for i := 0; i < n+1; i++ {
		fast = fast.Next
	}

	slow := dummyNode
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next

	return dummyNode.Next
}

// middleNode -----------------------------
func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// makeList -----------------------------
func makeList(arr []int) *ListNode {
	var head, tail *ListNode
	for _, v := range arr {
		newNode := &ListNode{Val: v}

		if head == nil {
			head = newNode
			tail = newNode
		} else {
			tail.Next = newNode
			tail = newNode
		}
	}
	return head
}
func (l *ListNode) Print() {
	curr := l
	for curr != nil {
		fmt.Println(curr.Val)
		curr = curr.Next
	}
	fmt.Println("end")
}
