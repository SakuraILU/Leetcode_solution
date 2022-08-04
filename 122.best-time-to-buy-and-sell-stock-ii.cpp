/*
 * @lc app=leetcode id=122 lang=cpp
 *
 * [122] Best Time to Buy and Sell Stock II
 */

// @lc code=start
#include<vector>
#include<algorithm>
using namespace std;
class Solution {
public:
    int maxProfit(vector<int>& prices) {
        vector<vector<int>> max_sum(prices.size(),vector<int>(2,0));
        //at the end of the day,if not have the stock, earn $0 of course
        // but if have the stock, the only choice is to sell it and earn prices[end_day]
        max_sum[prices.size()-1][1] = prices[prices.size()-1];
        for(int day = prices.size()-2; day >= 0; day--){
            for(int has = 1; has >=0; has--){
                max_sum[day][has] = max(
                {
                    //if not has the stock,policy can be buying the stock
                    (!has)?max_sum[day+1][!has]-prices[day]:0,
                    //if has the stock,policy can be selling the stock
                    (has)?max_sum[day+1][!has] + prices[day]:0,
                    //always can do nothing and observe the stock
                    max_sum[day+1][has]
                }
                );
            }
        }
        return max_sum[0][0];
    }
};
// @lc code=end

