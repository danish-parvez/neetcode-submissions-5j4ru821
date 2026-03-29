func majorityElement(nums []int) int {
    n := len(nums)
    for _, num := range nums {
        count := 0
        for _, i := range nums {
            if i == num {
                count++
            }
        }
        if count > n/2 {
            return num
        }
    }
    return -1
}
