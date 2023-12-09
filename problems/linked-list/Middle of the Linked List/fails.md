<- тут все решения, которые были зафейлены (ДНЕВНИК ОШИБОК)

```
 
func middleNode(head *ListNode) *ListNode {
   slow,fast:=head,head
   
   for fast!=nil|| fast.Next!=nil{
       slow=slow.Next
       fast=fast.Next.Next
   } 
   return slow
}
```
for fast!=nil || fast.Next!=nil
так как использовал || когда fast.Next!=nil код сработает и
обратится к fast.Next.Next что является nil и произойдет ошибка

**Используй &&**