/*
 * @lc app=leetcode id=304 lang=cpp
 *
 * [304] Range Sum Query 2D - Immutable
 */

// @lc code=start
#include <vector>
#include <iostream>

using namespace std;

#define RES (*ul_res)
class NumMatrix
{
private:
    vector<vector<int>> *ul_res;

public:
    NumMatrix(vector<vector<int>> &matrix)
    {
        int rows = matrix.size();
        int cols = matrix[0].size();
        ul_res = new vector<vector<int>>(rows, vector<int>(cols, 0));
        for (int row = 0; row < rows; row++)
        {
            for (int col = 0; col < cols; col++)
            {
                if (row == 0 && col == 0)
                    RES[row][col] = matrix[row][col];
                else if (row > 0 && col == 0)
                    RES[row][col] = matrix[row][col] + RES[row - 1][col];
                else if (row == 0 && col > 0)
                    RES[row][col] = matrix[row][col] + RES[row][col - 1];
                else
                    RES[row][col] = matrix[row][col] + RES[row][col - 1] + RES[row - 1][col] - RES[row - 1][col - 1];
            }
        }
    }

    int sumRegion(int row1, int col1, int row2, int col2)
    {
        if (row1 == 0 && col1 == 0)
            return RES[row2][col2];
        else if (row1 == 0)
            return RES[row2][col2] - RES[row2][col1 - 1];
        else if (col1 == 0)
            return RES[row2][col2] - RES[row1 - 1][col2];
        else
            return RES[row2][col2] - RES[row1 - 1][col2] - RES[row2][col1 - 1] + RES[row1 - 1][col1 - 1];
    }
};

/**
 * Your NumMatrix object will be instantiated and called as such:
 * NumMatrix* obj = new NumMatrix(matrix);
 * int param_1 = obj->sumRegion(row1,col1,row2,col2);
 */
// @lc code=end
