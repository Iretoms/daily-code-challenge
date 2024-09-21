package challenges

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	
    head := &ListNode{}
    current := head

    for list1 != nil && list2 != nil {
        if list1.Val <= list2.Val {
            current.Next = list1  
            list1 = list1.Next   
        } else {
            current.Next = list2  
            list2 = list2.Next    
        }
        current = current.Next  
    }

    if list1 != nil {
        current.Next = list1
    } else {
        current.Next = list2
    }

    return head.Next
}

