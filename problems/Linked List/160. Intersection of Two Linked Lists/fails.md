```
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    p1:=revGetIntersectionNode(headA)
    p2:=revGetIntersectionNode(headB)
var res *ListNode = nil

    for p1==p2{
res=p1
p1=p1.Next
p2=p2.Next
    }
return res
}

func revGetIntersectionNode(head *ListNode) *ListNode{
    var prew *ListNode= nil
cur:=head
tmp:=cur

for cur!=nil{
    tmp=cur
    cur=cur.Next
    tmp.Next=prew
    prew=tmp
}
return tmp
}
```



No intersection, ERROR: linked structure was modified.
когда сделал реверс, то изменил указатели


```
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
```
