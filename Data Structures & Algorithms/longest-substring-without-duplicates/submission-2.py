class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        l, length = 0, 0
        window = {}
        for r in range(len(s)):
            if s[r] in window:
                l = max(window[s[r]] + 1, l)
            window[s[r]] = r
            length = max(length, r - l + 1)
        return length