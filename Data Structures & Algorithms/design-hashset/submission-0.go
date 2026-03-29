type Node struct {
    key int
    next *Node
}

func NewNode(key int) *Node {
    return &Node{
        key: key,
        next: nil,
    }
}

type MyHashSet struct {
    set []*Node
}

func Constructor() MyHashSet {
    set := make([]*Node, 10000)
    for i := range set {
        set[i] = NewNode(0)
    }
    return MyHashSet{
        set: set,
    }
}

func (this *MyHashSet) Add(key int) {
    index := key % len(this.set)
    curr := this.set[index]
    for curr.next != nil {
        if curr.next.key == key {
            return
        }
        curr = curr.next
    }
    curr.next = NewNode(key)
}

func (this *MyHashSet) Remove(key int) {
    index := key % len(this.set)
    curr := this.set[index]
    for curr.next != nil {
        if curr.next.key == key {
            curr.next = curr.next.next
            return
        }
        curr = curr.next
    }
}

func (this *MyHashSet) Contains(key int) bool {
    index := key % len(this.set)
    curr := this.set[index]
    for curr.next != nil {
        if curr.next.key == key {
            return true
        }
        curr = curr.next
    }
    return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
 