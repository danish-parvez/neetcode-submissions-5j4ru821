type ListNode struct {
	key, val int
	next *ListNode
}

type MyHashMap struct {
	data []*ListNode
}

func Constructor() MyHashMap {
    data := make([]*ListNode, 1000)
	for i := range data {
		data[i] = &ListNode{key: -1, val: -1}
	}
	return MyHashMap{data: data}
}

func (this *MyHashMap) hash(key int) int {
	return key % len(this.data)
}

func (this *MyHashMap) Put(key int, value int) {
    curr := this.data[this.hash(key)]
	for curr.next != nil {
		if curr.next.key == key {
			curr.next.val = value
			return
		}
		curr = curr.next
	}
	curr.next = &ListNode{key: key, val: value}
}

func (this *MyHashMap) Get(key int) int {
    curr := this.data[this.hash(key)].next
	for curr != nil {
		if curr.key == key {
			return curr.val
		}
		curr = curr.next
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
   curr := this.data[this.hash(key)]
   for curr.next != nil {
		if curr.next.key == key {
			curr.next = curr.next.next
			return
		}
		curr = curr.next
   } 
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */