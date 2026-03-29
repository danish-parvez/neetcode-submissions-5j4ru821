class Solution:
    def mergeAlternately(self, word1: str, word2: str) -> str:
        l, r = 0, 0
        s = ""
        while l < len(word1) and r < len(word2):
            s += word1[l] + word2[r]
            l += 1
            r += 1
        if len(word1) < len(word2):
            s += word2[r:]
        elif len(word2) < len(word1):
            s += word1[l: ]
        return s