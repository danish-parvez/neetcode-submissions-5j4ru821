func twoSum(nums []int, target int) []int {
    indices := make(map[int]int)
    for i, num := range nums {
        indices[num] = i
    }
    for i, num := range nums {
        complement := target - num
        if j, ok := indices[complement]; ok && j != i {
            return []int{i, j}
        }
    }
    return []int{}
}
