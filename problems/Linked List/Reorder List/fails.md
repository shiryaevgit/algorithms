```
func reorderList(head *ListNode) {
	tmp := preMid(head)
	//				3 -> 4 -> 5
	// 	  		   tmp
	p2 := reversReorder(tmp.Next) // 4 -> 5
	//
	//								 5 -> 4
	//							     p2
	//
	tmp.Next = nil
	//			5 ->	 4	  -> nil
	//	 	 	  	  tmp.Next
	//

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
```

решение верное не понимаю. пройти по коду 