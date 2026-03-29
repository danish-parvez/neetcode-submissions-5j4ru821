type DynamicArray struct {
    length int
    capacity int
    arr []int
}

func NewDynamicArray(capacity int) *DynamicArray {
    return &DynamicArray{
        length: 0,
        capacity: capacity,
        arr: make([]int, capacity),
    }
}

func (da *DynamicArray) Get(i int) int {
    if i < da.length {
        return da.arr[i]
    }
    return 0
}

func (da *DynamicArray) Set(i int, n int) {
    if i < da.length {
        da.arr[i] = n
    }
}

func (da *DynamicArray) Pushback(n int) {
    if da.length == da.capacity {
        da.resize()
    }
    da.arr[da.length] = n
    da.length++
}

func (da *DynamicArray) Popback() int {
    if da.length > 0 {
        n := da.arr[da.length - 1]
        da.length--
        return n
    }
    return 0
}

func (da *DynamicArray) resize() {
    if da.length == da.capacity {
        da.capacity = 2 * da.capacity
        arrNew := make([]int, da.capacity)
        copy(arrNew, da.arr)
        da.arr = arrNew
    }
}

func (da *DynamicArray) GetSize() int {
    return da.length
}

func (da *DynamicArray) GetCapacity() int {
    return da.capacity
}
