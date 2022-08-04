/*
 * @lc app=leetcode id=650 lang=cpp
 *
 * [650] 2 Keys Keyboard
 */

// @lc code=start
#include<vector>
#include<algorithm>
#include<limits.h>
using namespace std;
class Solution {
private:
    vector<vector<int>> memo;
public:
    int minSteps(int n) {
        memo = vector<vector<int>>(n,vector<int>(n+1,-1));
        return dp(n,0,0);
    }

    int dp(int n, int pos,int cache_num){
        if(pos == n-1){memo[pos][cache_num] = 0; return 0;}
        
        memo[pos][cache_num] = min(
                    {
                    (cache_num > 0 && pos + cache_num < n)?
                        (memo[pos+cache_num][cache_num] != -1)?
                            memo[pos+cache_num][cache_num]
                            :dp(n,pos+cache_num,cache_num)
                        :INT_MAX,
                    (cache_num != pos+1)?
                        ((memo[pos][pos+1] != -1)?
                            memo[pos][pos+1]
                            :dp(n,pos,pos+1))
                        :INT_MAX
                    }
                );
                
        if(memo[pos][cache_num] != INT_MAX) memo[pos][cache_num] += 1;
        
        return memo[pos][cache_num];
    }
};
// @lc code=end

