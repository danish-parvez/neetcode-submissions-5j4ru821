class Solution:
    def climbStairs(self, n: int) -> int:
        memo = {}
        
        def recur_climb_stairs(n):
            if n <= 1:
                return 1
            if n in memo:
                return memo[n]
            memo[n] = recur_climb_stairs(n - 1) + recur_climb_stairs(n - 2)
            return memo[n]
        
        return recur_climb_stairs(n)