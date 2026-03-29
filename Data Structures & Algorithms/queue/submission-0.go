type ListNode struct {
    val int
    next *ListNode
}

type Deque struct {
    left *ListNode
    right *ListNode
}

func NewListNode(value int) *ListNode {
    return &ListNode{
        val: value,
        next: nil,
    }
}

func NewDeque() *Deque {
    return &Deque{
        left: nil,
        right: nil,
    }
}

func (d *Deque) IsEmpty() bool {
    if d.right == nil {
        return true
    }
    return false
}

func (d *Deque) Append(value int) {
    newNode := NewListNode(value)
    if d.IsEmpty() {
        d.left = newNode
        d.right = newNode
    } else {
        d.right.next = newNode
        d.right = d.right.next
    }
}

func (d *Deque) AppendLeft(value int) {
    newNode := NewListNode(value)
    if d.IsEmpty() {
        d.left = newNode
        d.right = newNode
    } else {
        newNode.next = d.left
        d.left = newNode
    }   
}

func (d *Deque) Pop() int {
    if d.IsEmpty() {
        return -1
    } else if d.left == d.right {
        val := d.left.val
        d.left = nil
        d.right = nil
        return val
    }
    curr := d.left
    for curr.next != nil && curr.next != d.right {
        curr = curr.next
    }
    val := curr.next.val
    curr.next = nil
    d.right = curr
    return val
}

func (d *Deque) PopLeft() int {
    if d.right == nil {
        return -1
    } else if d.left == d.right {
        val := d.left.val
        d.left = nil
        d.right = nil
        return val
    }
    val := d.left.val
    d.left = d.left.next
    return val
}
