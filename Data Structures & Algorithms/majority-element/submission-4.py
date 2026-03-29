class Solution:
    def majorityElement(self, nums: List[int]) -> int:
        counts = {}
        for num in nums:
            if num in counts:
                counts[num] += 1
            else:
                counts[num] = 1
        majority = len(nums) // 2
        for count in counts:
            if counts[count] > majority:
                return count
            