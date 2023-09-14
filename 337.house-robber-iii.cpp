/*
 * @lc app=leetcode id=337 lang=cpp
 *
 * [337] House Robber III
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */

// struct TreeNode
// {
//     int val;
//     TreeNode *left;
//     TreeNode *right;
//     TreeNode() : val(0), left(nullptr), right(nullptr) {}
//     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
//     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
// };

#include <unordered_map>
using namespace std;

#define MAX(a, b) ((a > b) ? a : b)

class Solution
{
private:
    unordered_map<TreeNode *, int> memo;

public:
    Solution() : memo(unordered_map<TreeNode *, int>()) {}

    int rob(TreeNode *root)
    {
        return dp(root);
    }

    int dp(TreeNode *node)
    {
        if (node == nullptr)
        {
            return 0;
        }

        if (memo.count(node) > 0)
        {
            return memo[node];
        }

        int money_ns, money_s = 0;
        // not stole
        money_ns = dp(node->left) + dp(node->right);

        // stole
        money_s += node->val;
        if (node->left != nullptr)
        {
            money_s += dp(node->left->left) + dp(node->left->right);
        }
        if (node->right != nullptr)
        {
            money_s += dp(node->right->left) + dp(node->right->right);
        }

        memo[node] = MAX(money_ns, money_s);

        return memo[node];
    }
};
// @lc code=end
