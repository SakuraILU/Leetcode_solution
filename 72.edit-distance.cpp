/*
 * @lc app=leetcode id=72 lang=cpp
 *
 * [72] Edit Distance
 */

// @lc code=start
#include<vector>
#include<iostream>
#include<algorithm>
#include<limits.h>
using namespace std;

class Solution {
public:
    int minDistance(string word1, string word2) {
        vector<vector<int>> edit_distance(word1.size()+1,vector<int>(word2.size()+1,0));
        //insert ed[i,j] = ed[i,j+1] + 1
        //replace ed[i,j] = ed[i+1,j+1] + 1
        //remove ed[i,j] = ed[i+1,j] + 1
        //do nothing ed[i,j] = ed[i+1,j+1]
        for(int i = word1.size(); i >= 0; i--){
            for(int j = word2.size(); j >=0; j--){
                
                if(i == word1.size() && j == word2.size())continue;
                
                edit_distance[i][j] = min(
                    {
                    (j < word2.size())?(edit_distance[i][j+1])+1:INT_MAX,
                    (i < word1.size() && j < word2.size())?edit_distance[i+1][j+1]+1:INT_MAX,
                    (i < word1.size())?edit_distance[i+1][j]+1:INT_MAX,
                    (i < word1.size() && j < word2.size() && (word1[i] == word2[j]))?edit_distance[i+1][j+1]:INT_MAX
                    }
                );
            }
        }
        return edit_distance[0][0];
    }
};
// @lc code=end

