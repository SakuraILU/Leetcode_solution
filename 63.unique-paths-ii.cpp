/*
 * @lc app=leetcode id=63 lang=cpp
 *
 * [63] Unique Paths II
 */

// @lc code=start
#include<vector>
#include<iostream>
using namespace std;

class Solution {
public:
    int uniquePathsWithObstacles(vector<vector<int>>& obstacleGrid) {
        vector<vector<unsigned int>> path_num(obstacleGrid.size(),vector<unsigned int>(obstacleGrid[0].size(),0));
        
        if(obstacleGrid[obstacleGrid.size()-1][obstacleGrid[0].size()-1] == 0) 
            path_num[obstacleGrid.size()-1][obstacleGrid[0].size()-1] = 1;

        for(int i = obstacleGrid.size()-1;i >= 0; i--){
            for(int j = obstacleGrid[0].size()-1; j >= 0; j--){
                if(i == obstacleGrid.size()-1 && j == obstacleGrid[0].size()-1) continue;
                
                if(obstacleGrid[i][j] == 1){ path_num[i][j] = 0;continue;}
              
                path_num[i][j] = (   ( (j < obstacleGrid[0].size()-1 && obstacleGrid[i][j+1] == 0)?path_num[i][j+1]:0 ) 
                                   + 
                                     ( (i < obstacleGrid.size()-1 && obstacleGrid[i+1][j] == 0)?path_num[i+1][j]:0 )
                                 );
            }
        }
        return path_num[0][0];
    }

    // void print(vector<vector<int>> &path_num){
    //     for (size_t i = 0; i < path_num.size(); i++)
    //     {
    //         for (size_t j = 0; j < path_num[0].size(); j++)
    //         {
    //             cout<<path_num[i][j]<<'|';
    //         }
    //         cout<<endl;
    //         /* code */
    //     }
    //     cout<<endl<<"============="<<endl;
    // }
};
// @lc code=end

