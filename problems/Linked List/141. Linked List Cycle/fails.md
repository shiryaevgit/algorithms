```
 
 for fast!=nil{ // не прописал fast=fast.Next
     slow=slow.Next
     fast=fast.Next.Next
 
     if slow==fast{
         return true
     }
 }   
 return false
}
```

```
func hasCycle(head *ListNode) bool {
slow,fast:=head,head

for fast!=nil || fast.Next!=nil{ //  fast!=nil && fast.Next!=nil
slow=slow.Next
fast=fast.Next.Next

     if slow==fast{
         return true
     }
}   
return false
}
```
fast!=nil && fast.Next!=nil
логическое И потому, что нам нужно, чтобы выполнялись ОБА условия
```
for fast!=nil && fast.Next!=nil
```
fast != nil (чтобы не выйти за границы списка)
fast.Next!=nil (чтобы можно было перейти на следующий элемент)