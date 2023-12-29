
```
func removeNthFromEnd(head *ListNode, n int) *ListNode {
lengtn:=0
curr:=head

    for curr!=nill{
        length++
        curr=curr.next
    }
    
    dummyNode:=&ListNode{next:head}
    curr=dummyNode
    
    for i:=0;i<length-n-1;i++{
        curr=curr.next
    }
    curr=curr.next.next
    
    return dummyNode.next
}
```
