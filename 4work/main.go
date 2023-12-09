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
	a := []int{1, 2, 3, 4, 5, 6}
	l1 := makeList(a)

	r := rotateRight(l1, 1)

	printList(r)

}
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head // Возвращаем исходный список, если он пуст или не требует сдвига
	}

	length := 1
	cur := head

	for cur.Next != nil {
		cur = cur.Next
		length++
	}
	cur.Next = head // Делаем цикл

	k = k % length // вернет остаток, то насколько надо сдвинуть
	stepsToNewHead := length - k

	for i := 0; i < stepsToNewHead; i++ {
		cur = cur.Next
	}

	newHead := cur.Next
	cur.Next = nil // рвем цикл перед newHead

	return newHead
}

func nextLargerNodes(head *ListNode) []int {
	cur, tmp := head, head
	length := 0
	i := 0

	for tmp != nil {
		length++
		tmp = tmp.Next
	}
	res := make([]int, length)

	for cur != nil {
		tmp = cur.Next

		for tmp != nil {

			if cur.Val < tmp.Val {
				res[i] = tmp.Val
				break
			} else {
				tmp = tmp.Next
			}
		}
		i++
		cur = cur.Next
	}
	return res
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			break
		}
	}

	if fast == nil || fast.Next == nil {
		return nil
	}

	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}

	return fast
}

func hasCycle(head *ListNode) bool {

	slow, fast := head, head

	for fast != nil || fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}
	return false
}

func mergeNLists(lists []*ListNode) *ListNode {
	dummyNode := &ListNode{}
	cur := dummyNode
	for {
		minNode, indexList := MinOfLists(lists)
		if minNode == nil {
			break
		}
		cur.Next = minNode
		lists[indexList] = lists[indexList].Next
		cur = cur.Next
	}

	return dummyNode.Next
}

func MinOfLists(lists []*ListNode) (node *ListNode, index int) {
	min := math.MaxInt32

	for i, l := range lists {
		if l == nil {
			continue
		}
		if l.Val < min {
			node = l
			index = i
			min = l.Val
		}
	}
	return
}

func merge2Lists(l1, l2 *ListNode) *ListNode {
	dummyNode := &ListNode{}
	cur := dummyNode

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next

		if l1 != nil {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
	}
	return dummyNode.Next
}

func deleteN(head *ListNode, n int) *ListNode {
	dummyNode := &ListNode{Next: head}
	cur := dummyNode

	for i := 0; i < n-1; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next

	return dummyNode.Next
}

func deleteNEnd2(head *ListNode, n int) *ListNode {
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

func deleteNEnd(head *ListNode, n int) *ListNode {
	dummyNode := &ListNode{Next: head}
	cur := dummyNode
	length := 0

	for cur != nil {
		length++
		cur = cur.Next
	}

	cur = dummyNode
	for i := 0; i < length-n-1; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next

	return dummyNode.Next
}

func makeList(a []int) *ListNode {
	var head, tail *ListNode

	for _, v := range a {
		newNode := &ListNode{Val: v}

		if head == nil {
			head, tail = newNode, newNode
		} else {
			tail.Next = newNode
			tail = newNode
		}
	}
	return head
}

func revers(head *ListNode) *ListNode {
	var prev *ListNode = nil
	cur, tmp := head, head

	for cur != nil {
		tmp = cur
		cur = cur.Next
		tmp.Next = prev
		prev = tmp
	}
	return tmp
}

func mid(head *ListNode) *ListNode {
	fast, slow := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
