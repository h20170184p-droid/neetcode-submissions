type Node struct {
    val  int
    prev *Node 
    next *Node 
}

type MyLinkedList struct {
    head   *Node
    tail   *Node
    length int
}

func Constructor() MyLinkedList {
    dummyHead := &Node{val: -1}
    dummyTail := &Node{val: -1}

    dummyHead.next = dummyTail
    dummyTail.prev = dummyHead
    return MyLinkedList{
        head:   dummyHead,
        tail:   dummyTail,
        length: 0,
    }
}

func (this *MyLinkedList) Get(index int) int {
    if index < 0 || index >= this.length {
        return -1
    }

    if index < this.length/2 {
        presentAddress := this.head.next
        for range index { // Removed 'i' to fix compiler error
            presentAddress = presentAddress.next
        }
        return presentAddress.val
    } else {
        presentAddress := this.tail.prev
        backwardSteps := (this.length - 1) - index
        for range backwardSteps { // Removed 'i' to fix compiler error
            presentAddress = presentAddress.prev
        }
        return presentAddress.val
    }
}

func (this *MyLinkedList) AddAtHead(val int)  {
    newNodeAddress := &Node{
        val:  val,
        prev: this.head,
        next: this.head.next,
    }
    this.head.next.prev = newNodeAddress
    this.head.next = newNodeAddress
    this.length += 1
}

func (this *MyLinkedList) AddAtTail(val int)  {
    newNodeAddress := &Node {
        val:  val,
        prev: this.tail.prev,
        next: this.tail,
    }
    this.tail.prev.next = newNodeAddress
    this.tail.prev = newNodeAddress
    this.length += 1
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
    if index < 0 || index > this.length {
        return
    }
    if index == 0 {
        this.AddAtHead(val) // Updates length internally
        return
    }
    if index == this.length {
        this.AddAtTail(val) // Updates length internally
        return
    }

    var targetAddress *Node
    if index < this.length/2 {
        targetAddress = this.head.next
        for range index { // Removed 'i'
            targetAddress = targetAddress.next
        }
    } else {
        targetAddress = this.tail.prev
        backwardSteps := (this.length - 1) - index
        for range backwardSteps { // Removed 'i'
            targetAddress = targetAddress.prev
        }
    }

    newNodeAddress := &Node{
        val:  val,
        prev: targetAddress.prev,
        next: targetAddress,
    }

    targetAddress.prev.next = newNodeAddress 
    targetAddress.prev = newNodeAddress       

    this.length++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
    if index < 0 || index >= this.length {
        return
    }

    var nodeToDeleteAddress *Node
    if index < this.length/2 {
        nodeToDeleteAddress = this.head.next
        for range index { // Removed 'i'
            nodeToDeleteAddress = nodeToDeleteAddress.next
        }
    } else {
        nodeToDeleteAddress = this.tail.prev
        backwardSteps := (this.length - 1) - index
        for range backwardSteps { // Removed 'i'
            nodeToDeleteAddress = nodeToDeleteAddress.prev
        }
    }

    leftNeighbor := nodeToDeleteAddress.prev
    rightNeighbor := nodeToDeleteAddress.next

    leftNeighbor.next = rightNeighbor
    rightNeighbor.prev = leftNeighbor

    this.length--
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */