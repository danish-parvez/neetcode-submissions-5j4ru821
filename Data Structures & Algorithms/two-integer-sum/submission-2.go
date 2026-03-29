func twoSum(nums []int, target int) []int {
    prevMap := make(map[int]int)
    for i, num := range nums {
        comp := target - num
        if j, ok := prevMap[comp]; ok {
            return []int{j, i}
        }
        prevMap[num] = i
    }
    return []int{}
}
