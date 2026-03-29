type ListNode struct {
	key int
	val int
	next *ListNode
}

func NewListNode(key, value int) *ListNode {
	return &ListNode {
		key: key,
		val: value,
		next: nil,
	}
}

type HashTable struct {
	size int
	capacity int
	table []*ListNode
}

func NewHashTable(capacity int) *HashTable {
	return &HashTable{
		size: 0,
		capacity: capacity,
		table: make([]*ListNode, capacity),
	}
}

func (ht *HashTable) hash(key int) int {
	return key % ht.capacity
}

func (ht *HashTable) Insert(key, value int) {
	index := ht.hash(key)
	node := ht.table[index]
	if node == nil {
		ht.table[index] = NewListNode(key, value)
		ht.size++
	} else {
		var prev *ListNode
		for node != nil {
			if node.key == key {
				node.val = value
				return
			}
			prev, node = node, node.next
		}
		prev.next = NewListNode(key, value)
		ht.size++
	}
	if float64(ht.size)/float64(ht.capacity) >= 0.5 {
		ht.Resize()
	}
}

func (ht *HashTable) Get(key int) int {
	index := ht.hash(key)
	node := ht.table[index]
	for node != nil {
		if node.key == key {
			return node.val
		}
		node = node.next
	}
	return -1
}

func (ht *HashTable) Remove(key int) bool {
	index := ht.hash(key)
	node := ht.table[index]
	var prev *ListNode
	for node != nil {
		if node.key == key {
			if prev != nil {
				prev.next = node.next
			} else {
				ht.table[index] = node.next
			}
			ht.size--
			return true
		}
		prev, node = node, node.next
	}
	return false
}

func (ht *HashTable) GetSize() int {
	return ht.size
}

func (ht *HashTable) GetCapacity() int {
	return ht.capacity
}

func (ht *HashTable) Resize() {
	newCapacity := ht.capacity * 2
	newTable := make([]*ListNode, newCapacity)
	for _, node := range ht.table {
		for node != nil {
			index := node.key % newCapacity
			if newTable[index] == nil {
				newTable[index] = NewListNode(node.key, node.val)
			} else {
				newNode := newTable[index]
				for newNode.next != nil {
					newNode = newNode.next
				}
				newNode.next = NewListNode(node.key, node.val)
			}
			node = node.next
		}
	}
	ht.capacity = newCapacity
	ht.table = newTable
}
