/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 * Val  int
 * Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    // 1. Create a fake dummy node to start our new list
    dummy := &ListNode{Val: -1}
    
    // 'tail' will be the moving pointer that updates our new list
    tail := dummy 

    // 2. Loop as long as BOTH lists have nodes left to compare
    for list1 != nil && list2 != nil {
        if list1.Val <= list2.Val {
            // Your logic: list1 is smaller (or equal), attach it!
            tail.Next = list1
            list1 = list1.Next // "Remove" the head of list1 by moving it forward
        } else {
            // list2 is smaller, attach it!
            tail.Next = list2
            list2 = list2.Next // "Remove" the head of list2 by moving it forward
        }
        
        // Move our merged list's tail forward to the node we just attached
        tail = tail.Next
    }

    // 3. Clean up the leftover nodes
    // If one list finishes early, just hook the rest of the other list to the end
    if list1 != nil {
        tail.Next = list1
    } else if list2 != nil {
        tail.Next = list2
    }

    // 4. Return the actual start of the sorted list (skipping our fake dummy node)
    return dummy.Next
}
