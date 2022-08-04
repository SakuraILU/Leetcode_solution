/*
 * @lc app=leetcode id=309 lang=cpp
 *
 * [309] Best Time to Buy and Sell Stock with Cooldown
 */

// @lc code=start
#include<vector>
#include<algorithm>
using namespace std;

class Solution {
public:
    int maxProfit(vector<int>& prices) {
        vector<vector<int>> max_sum(prices.size()+1,vector<int>(2,0));
        
        max_sum[prices.size()-1][1] = prices[prices.size()-1];
        
        for(int day = prices.size()-2; day >= 0; day--){
            for(int has = 1; has >= 0; has--){
                max_sum[day][has] = max(
                    {
                        (!has)?max_sum[day+1][!has] - prices[day]:0,
                        (has)?max_sum[day+2][!has] + prices[day]:0,
                        max_sum[day+1][has]
                    }
                );
            }
        }
        return max_sum[0][0];
    }
};
// @lc code=end

