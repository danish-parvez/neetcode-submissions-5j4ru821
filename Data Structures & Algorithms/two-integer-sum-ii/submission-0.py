class Solution:
    def twoSum(self, numbers: List[int], target: int) -> List[int]:
        l = 0
        r = len(numbers) - 1
        res = []
        while l < r:
            total = numbers[l] + numbers[r]
            if total == target:
                res = [l + 1, r + 1]
                return res
            
            if total > target:
                r -= 1
            else:
                l += 1
        
        return res