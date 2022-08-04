/*
 * @lc app=leetcode id=213 lang=cpp
 *
 * [213] House Robber II
 */

// @lc code=start
#include<vector>
#include<iostream>
#include<algorithm>
using namespace std;

class Solution {
public:
    int rob(vector<int>& nums) {
        if(nums.size() == 1) return nums[0];
        if(nums.size() == 2) return max(nums[0],nums[1]);

        return max(  
                     nums[0] + ( 
                                (nums.size()>=4)?rob_no_circle(  vector<int>(nums.begin()+2,nums.end()-1)  ):0 
                               ),
                     (nums.size()>=2)? rob_no_circle( vector<int>(nums.begin()+1,nums.end()) ):0 
                  );
    }
    
    int rob_no_circle(const vector<int>& nums){
        vector<unsigned> money(nums.size()+2,0);
        for(int i = nums.size()-1; i >= 0; i--){
            money[i] = max(money[i+2]+nums[i],money[i+1]);
        }
        return money[0];
    }
};
// @lc code=end

