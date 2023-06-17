/*
 * @lc app=leetcode id=224 lang=cpp
 *
 * [224] Basic Calculator
 */

// @lc code=start
#include <iostream>
#include <string>
#include <algorithm>
using namespace std;

class Solution
{
private:
    bool is_unary_op(string &s, int lo, int hi, int op_pos)
    {
        for (int i = op_pos - 1; i >= lo; --i)
        {
            if (s[i] == ' ')
                continue;
            if (s[i] == '(' || s[i] == '-' || s[i] == '+')
                return true;
            else
                return false;
        }
        return true;
    }

    int add_exp(string &s, int lo, int hi)
    {
        int last_op = -1;
        int depth = 0;
        for (int i = hi - 1; i >= lo; --i)
        {
            if (s[i] == ')')
                depth++;
            else if (s[i] == '(')
                depth--;

            if (depth == 0)
                if (s[i] == '+' || s[i] == '-')
                {
                    if (is_unary_op(s, lo, hi, i))
                        continue;

                    last_op = i;
                    // cout << "op " << s[last_op] << std::endl;
                    break;
                }
        }

        if (last_op == -1)
            return unary_exp(s, lo, hi);
        else
        {
            if (s[last_op] == '+')
                return add_exp(s, lo, last_op) + unary_exp(s, last_op + 1, hi);
            else
                return add_exp(s, lo, last_op) - unary_exp(s, last_op + 1, hi);
        }
    }

    int unary_exp(string &s, int lo, int hi)
    {
        int first_op = -1;
        int depth = 0;
        for (int i = lo; i < hi; ++i)
        {
            if (s[i] == ')')
                depth++;
            else if (s[i] == '(')
                depth--;

            if (depth == 0)
                if (s[i] == '+' || s[i] == '-')
                {
                    first_op = i;
                    // cout << "op " << s[last_op] << std::endl;
                    break;
                }
        }

        if (first_op == -1)
            return primary_exp(s, lo, hi);
        else
        {
            if (s[first_op] == '+')
                return unary_exp(s, first_op + 1, hi);
            else
                return -unary_exp(s, first_op + 1, hi);
        }
    }

    int primary_exp(string &s, int lo, int hi)
    {
        if (s[lo] == '(' && s[hi - 1] == ')')
            return add_exp(s, lo + 1, hi - 1);
        else
            return lval(s, lo, hi);
    }

    int lval(string &s, int lo, int hi)
    {
        // cout << s.substr(lo, hi - lo) << endl;
        return stoi(s.substr(lo, hi - lo));
    }

public:
    int calculate(string s)
    {
        std::string::iterator end_pos = std::remove(s.begin(), s.end(), ' ');
        s.erase(end_pos, s.end());
        return add_exp(s, 0, s.size());
    }
};
// @lc code=end
