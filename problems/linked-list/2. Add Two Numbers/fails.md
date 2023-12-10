```

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1 = reverss(l1)
	l2 = reverss(l2)
	num := ListToNum(l1) + ListToNum(l2)
	arr := numToSlice(num)
	fmt.Println(arr)
	res := MakeList(arr)
    res = reverss(res)
	return res
}
func MakeList(arr []int) *ListNode {
var head *ListNode
	tail := head

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

func reverss(head *ListNode) *ListNode {
	var prew *ListNode = nil

	cur, tmp := head, head

	for cur != nil {
		tmp = cur
		cur = cur.Next

		tmp.Next = prew
		prew = tmp
	}
	return tmp
}

func ListToNum(head *ListNode) int {
	cur := head
	res := 0

	for cur != nil {
		res = res*10 + cur.Val
		cur = cur.Next
	}
	return res
}

func numToSlice(n int) []int {
	if n == 0 {
		return []int{0}
	}

	digits := []int{}

	for n > 0 {
		digit := n % 10
		digits = append([]int{digit}, digits...)
		n /= 10
	}
	return digits
}
```