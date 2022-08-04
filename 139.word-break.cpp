/*
 * @lc app=leetcode id=139 lang=cpp
 *
 * [139] Word Break
 */

// @lc code=start
#include<string>
#include<vector>
#include<iostream>
using namespace std;

class Solution {
public:
    bool wordBreak(string s, vector<string>& wordDict) {
        vector<bool> can_segment(s.size()+1,false);
        
        can_segment[s.size()] = true;

        for(int i = s.size()-1; i >= 0; i--){
            for(auto world:wordDict){
                if( ((s.size()-i) >= world.size()) && (s.substr(i,world.size()) == world) )
                    can_segment[i] = can_segment[i+world.size()] || can_segment[i];
            }
        }

        return can_segment[0]; 
    }
};
// @lc code=end

