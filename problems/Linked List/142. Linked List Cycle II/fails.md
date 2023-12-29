```
func detectCycle(head *ListNode) *ListNode {
    fast,slow:=head,head
    
    for fast!=nil && fast.Next!=nil{
        slow=slow.Next
        fast=fast.Next.Next
    
        if slow==fast{
            return slow
        } 
    }
    return nil
}
```
они сошлись где-то в списке, 
не обязательно на элементе, где начинается цикл

