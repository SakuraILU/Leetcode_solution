/*
 * @lc app=leetcode id=10 lang=cpp
 *
 * [10] Regular Expression Matching
 */

// @lc code=start
#include<string>
#include<vector>
using namespace std;

#define bool_or(a,b,c) ((a)||(b)||(c)) 
class Solution {
public:
    bool isMatch(string s, string p) {
        vector<vector<bool>> match(s.size()+1,vector<bool>(p.size()+1,false));
        match[s.size()][p.size()] = 1;
        for (int i = s.size(); i >= 0; i--)
        {
            for (int j = p.size(); j >= 0; j--)
            {
                if(i == s.size() && j == p.size()) continue;

                match[i][j] = bool_or(
                                       ((i < s.size() && j < p.size()) && ( p[j] == s[i] || p[j] == '.' ))?match[i+1][j+1]:false,
                                       ((i < s.size() && j < p.size()) && ( p[j+1] == '*' && (p[j]== s[i] || p[j] == '.') ))?match[i+1][j]:false,
                                       ((j < p.size()-1) && p[j+1] == '*')?match[i][j+2]:false
                                     );
            }
        }
        return match[0][0];
    }

    // void print(vector<vector<bool>> &match){
    //     for (size_t i = 0; i < match.size(); i++)
    //     {
    //         for (size_t j = 0; j < match[0].size(); j++)
    //         {
    //             cout<<match[i][j]<<'|';
    //         }
    //         cout<<endl;
    //         /* code */
    //     }
    //     cout<<endl<<"============="<<endl;
    // }
};
// @lc code=end

