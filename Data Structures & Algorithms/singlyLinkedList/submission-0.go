type Node struct {
    value       int
    nextAddress *Node
}

type List struct {
    head *Node
    tail *Node
}

func NewNode(val int, nextAddr *Node) *Node {
    return &Node{
        value:       val,
        nextAddress: nextAddr,
    }
}

func NewLinkedList() *List {
    startNode := NewNode(-1, nil) 
    return &List{
        head: startNode,
        tail: startNode,
    }
}

func (ll *List) Get(index int) int {
    presentNodeAddress := ll.head.nextAddress
    i := 0
    for presentNodeAddress != nil {
        if i == index {
            return presentNodeAddress.value
        }
        i++
        presentNodeAddress = presentNodeAddress.nextAddress
    }
    return -1
}

func (ll *List) InsertHead(val int) {
    newNode := NewNode(val, ll.head.nextAddress)
    ll.head.nextAddress = newNode
    if newNode.nextAddress == nil {
        ll.tail = newNode
    }
}

func (ll *List) InsertTail(val int) {
    ll.tail.nextAddress = NewNode(val, nil)
    ll.tail = ll.tail.nextAddress
}

func (ll *List) Remove(index int) bool {
    i := 0
    presentNodeAddress := ll.head
    
    // Strict bounds check: make sure we don't advance if we are already out of bounds
    for i < index && presentNodeAddress != nil {
        i++
        presentNodeAddress = presentNodeAddress.nextAddress
    }

    // Walked to the node that is just before i-th node
    if presentNodeAddress != nil && presentNodeAddress.nextAddress != nil {
        // If the node we are deleting happens to be the tail, 
        // move the tail pointer back to our present node
        if presentNodeAddress.nextAddress == ll.tail {
            ll.tail = presentNodeAddress
        }
        
        // Snip the node out by pointing over it
        presentNodeAddress.nextAddress = presentNodeAddress.nextAddress.nextAddress
        return true
    }
    
    return false
}

// Fixed the receiver name from LinkedList to List to match your struct
func (ll *List) GetValues() []int {
    var res []int
    presentNodeAddress := ll.head.nextAddress // Start after the dummy node
    
    for presentNodeAddress != nil {
        res = append(res, presentNodeAddress.value)
        presentNodeAddress = presentNodeAddress.nextAddress
    }
    
    return res
}