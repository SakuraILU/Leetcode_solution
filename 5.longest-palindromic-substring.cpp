/*
 * @lc app=leetcode id=5 lang=cpp
 *
 * [5] Longest Palindromic Substring
 */

// @lc code=start
#include<string>
#include<vector>
//#include<iostream>
using namespace std;
class Solution {
public:
    string longestPalindrome(string s) {
        vector<vector<bool>> isPalindromic(s.size(),vector<bool>(s.size(),true));

        int max_len = 1;
        int start_pos = 0;
        for (int i = 1; i < s.size(); i++)
        {
            for (int j = 0; j < s.size()-i; j++)
            {
                //[j][j+1]
                if(isPalindromic[j+1][j+i-1] && (s[j] == s[j+i])){
                    // isPalindromic[j][j+i] = true;
                    if(i+1 > max_len){
                        max_len = i+1;
                        start_pos = j;
                    }
                }
                else
                    isPalindromic[j][j+i] = false;
            }
        }

        return s.substr(start_pos,max_len);
    }

    // void print(vector<vector<bool>> &isPalindromic){
    //     for (size_t i = 0; i < isPalindromic.size(); i++)
    //     {
    //         for (size_t j = 0; j < isPalindromic.size(); j++)
    //         {
    //             cout<<isPalindromic[i][j]<<'|';
    //             /* code */
    //         }
    //         cout<<endl;
    //         /* code */
    //     }
    //     cout<<endl<<"============="<<endl;
    // }

};
// @lc code=end

