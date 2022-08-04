/*
 * @lc app=leetcode id=53 lang=cpp
 *
 * [53] Maximum Subarray
 */

// @lc code=start
#include<vector>
#include<limits.h>
using namespace std;
class Solution {
public:
    int maxSubArray(vector<int>& nums) {
        vector<int> max_sums(nums.size()+1,0);
        int max_sum = INT_MIN;
        for(int i = nums.size()-1; i >= 0; i--){
            //[i] -> [i+1] + cur
            max_sums[i] = ( (max_sums[i+1] > 0)?max_sums[i+1]:0 ) + nums[i];
            if(max_sums[i] > max_sum)
                max_sum = max_sums[i];
        }

        return max_sum;
    }
};
// @lc code=end

