/*
 * @lc app=leetcode id=1143 lang=cpp
 *
 * [1143] Longest Common Subsequence
 */

// @lc code=start
#include<vector>
#include<string>
#include<algorithm>
using namespace std;
class Solution {
public:
    int longestCommonSubsequence(string text1, string text2) {
        vector<vector<int>> longest_com_subseq(text1.size()+1,vector<int>(text2.size()+1,0));

        for(int i = text1.size(); i >= 0; i--){
            for(int j = text2.size(); j >= 0; j--){
                if(i == text1.size() && j == text2.size()) continue;

                longest_com_subseq[i][j] = max(
                    {
                        (i < text1.size() && j < text2.size() && text1[i] == text2[j])?longest_com_subseq[i+1][j+1]+1:0,//match [i,j]->[i+1,j+1] + 1
                        (i < text1.size())?longest_com_subseq[i+1][j]:0,//not match [i,j] -> [i+1,j]
                        (j < text2.size())?longest_com_subseq[i][j+1]:0 // or [i,j+1]
                    }
                );
            }
        }
        return longest_com_subseq[0][0];
    }
};
// @lc code=end

