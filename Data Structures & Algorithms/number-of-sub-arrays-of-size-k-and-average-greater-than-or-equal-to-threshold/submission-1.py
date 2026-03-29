class Solution:
    def numOfSubarrays(self, arr: List[int], k: int, threshold: int) -> int:
        window = []
        l = 0
        res = []
        for r in range(len(arr)):
            if r - l + 1 > k:
                window.remove(arr[l])
                l += 1
            window.append(arr[r])
            if len(window) == k and sum(window) // k >= threshold:
                res.append(window)
            
        return len(res)