```
func deleteDuplicates(head *ListNode) *ListNode {
	cur := head

	for cur != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		}
		cur = cur.Next
	}
	return head
}
```
panic: runtime error: invalid memory address or nil pointer dereference

&& cur.Next != nil

```
func deleteDuplicates(head *ListNode) *ListNode {
	cur := head

	for cur != nil && cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		}
		cur = cur.Next
	}
	return head
}
```
Input
head =
[1,1,1]

Use Testcase
Output
[1,1]
Expected
[1]



```
func deleteDuplicates(head *ListNode) *ListNode {
	cur := head
	tmp := cur.Next // если путой передал. то здесь будет ошибка!
	panic: runtime error: invalid memory address or nil pointer dereference


	for cur != nil && cur.Next != nil {

		if cur.Val == tmp.Val {
			cur.Next = cur.Next.Next
			tmp = cur.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}
```