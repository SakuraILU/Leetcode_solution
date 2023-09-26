#
# @lc app=leetcode id=887 lang=python3
#
# [887] Super Egg Drop
#

# @lc code=start
class Solution:
    def superEggDrop(self, k: int, n: int) -> int:
        memo = dict()
        def dp(K, N) -> int:
            # base case
            if K == 1: return N
            if N == 0: return 0
            # 避免重复计算
            if (K, N) in memo:
                return memo[(K, N)]

            res = float('INF')
            # 穷举所有可能的选择
            for i in range(1, N + 1):
                res = min(res, 
                        max(
                                dp(K, N - i), 
                                dp(K - 1, i - 1)
                            ) + 1
                    )
            # 记入备忘录
            memo[(K, N)] = res
            return res
        
        return dp(k, n)
# @lc code=end

