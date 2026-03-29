func majorityElement(nums []int) int {
    res := make(map[int]int)
    for _, num := range nums {
        res[num]++
    }
    for num, count := range res {
        if count >= len(nums)/2 {
            return num
        }
    }
    return 0
}
