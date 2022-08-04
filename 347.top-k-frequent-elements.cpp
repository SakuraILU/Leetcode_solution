/*
 * @lc app=leetcode id=347 lang=cpp
 *
 * [347] Top K Frequent Elements
 */

// @lc code=start
#include<vector>
using namespace std;
class Solution {
public:
    vector<int> topKFrequent(vector<int>& nums, int k) {
        
    }

private:
    int partition(vector<int>& nums){
        int pos = 0;
        for(int i = 0; i < nums.size(); ++i){
            if(nums[i] <= nums[pos]){
                for(int j = i; j > pos; --j)
                    swap(nums[j],nums[j-1]);
                ++pos;
            }
        }
        return pos;
    }
};
// @lc code=end

