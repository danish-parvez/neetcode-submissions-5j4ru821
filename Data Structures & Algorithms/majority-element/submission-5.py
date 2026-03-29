class Solution:
    def majorityElement(self, nums: List[int]) -> int:
        counts = {}
        res = maxCount = 0
        for num in nums:
            if num in counts:
                counts[num] += 1
            else:
                counts[num] = 1
            if maxCount < counts[num]:
                res = num
                maxCount = counts[num]
        return res
            