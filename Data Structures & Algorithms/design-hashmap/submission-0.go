type Pair struct {
    key int
    value int
}

func NewPair(key, val int) *Pair {
    return &Pair{
        key: key,
        value: val,
    }
}

type MyHashMap struct {
    size int
    capacity int
    hashMap []*Pair
}


func Constructor() MyHashMap {
    return MyHashMap{
        size: 0,
        capacity: 2,
        hashMap: make([]*Pair, 2),
    }
}

func (this *MyHashMap) Hash(key int) int {
    index := key % this.capacity
    return index
}

func (this *MyHashMap) Rehash() {
    this.capacity = 2 * this.capacity
    newMap := make([]*Pair, this.capacity)
    oldMap := this.hashMap
    this.hashMap = newMap
    this.size = 0
    for _, v := range oldMap {
        if v != nil {
            this.Put(v.key, v.value)
        }
    }
}


func (this *MyHashMap) Put(key int, value int)  {
    index := this.Hash(key)
    for {
        if this.hashMap[index] == nil {
            this.hashMap[index] = NewPair(key, value)
            this.size++
            if this.size >= this.capacity/2 {
                this.Rehash()
            }
            return
        } else if this.hashMap[index].key == key {
            this.hashMap[index].value = value
            return
        }
        index++
        index = index % this.capacity
    }
}


func (this *MyHashMap) Get(key int) int {
    index := this.Hash(key)
    for this.hashMap[index] != nil {
        if this.hashMap[index].key == key {
            return this.hashMap[index].value
        }
        index++
        index = index % this.capacity
    }
    return -1
}


func (this *MyHashMap) Remove(key int)  {
    if this.Get(key) == -1 {
        return
    }
    index := this.Hash(key)
    for {
        if this.hashMap[index].key == key {
            this.hashMap[index] = nil
            this.size--
            return
        }
        index++
        index = index % this.capacity
    }
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */