/*
 * @lc app=leetcode id=215 lang=cpp
 *
 * [215] Kth Largest Element in an Array
 */

// @lc code=start
#include<vector>
#include<iostream>
using namespace std;
class Solution {
public:
    int findKthLargest(vector<int>& nums, int k) {
        return findKthLargestNum(nums,0,nums.size(),k);
    }
private:
    int findKthLargestNum(vector<int>& nums, int low, int high, int k){
        int pos = partition(nums,low,high);
        int rank = high - pos;
        if (k > rank) return findKthLargestNum(nums,low,pos,k - rank);
        else if(k < rank) return findKthLargestNum(nums,pos+1,high,k);
        
        // if k == nums[pos]
        return nums[pos]; 
    }

    int partition(vector<int>& nums, int low, int high){
        int pos = low;
        for(int i = low + 1; i < high; ++i){
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

