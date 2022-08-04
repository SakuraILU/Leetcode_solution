/*
 * @lc app=leetcode id=198 lang=cpp
 *
 * [198] House Robber
 */

// @lc code=start
#include<vector>
#include<iostream>
using namespace std;
#define max(a,b) ((a>b)?a:b)
class Solution {
public:
    int rob(vector<int>& nums) {
        vector<unsigned> money(nums.size()+2,0);
        for(int i = nums.size()-1; i >= 0; i--){
            money[i] = max(money[i+2]+nums[i],money[i+1]);
        }
        return money[0];
    }
};
// @lc code=end

