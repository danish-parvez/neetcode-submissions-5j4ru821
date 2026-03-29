class Solution:
    def removeElement(self, nums: List[int], val: int) -> int:
        i = 0
        k = 0
        while i < len(nums) - k:
            if nums[i] == val:
                k += 1
                for j in range(i + 1, len(nums)):
                    nums[j-1] = nums[j]
                continue
            i += 1
        return len(nums) - k