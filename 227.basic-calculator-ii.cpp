/*
 * @lc app=leetcode id=227 lang=cpp
 *
 * [227] Basic Calculator II
 */

// @lc code=start

#include <iostream>
#include <string>
#include <algorithm>
using namespace std;

class Solution
{
private:
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
                    last_op = i;
                    // cout << "op " << s[last_op] << std::endl;
                    break;
                }
        }

        if (last_op == -1)
            return mul_exp(s, lo, hi);
        else
        {
            if (s[last_op] == '+')
                return add_exp(s, lo, last_op) + mul_exp(s, last_op + 1, hi);
            else
                return add_exp(s, lo, last_op) - mul_exp(s, last_op + 1, hi);
        }
    }

    int mul_exp(string &s, int lo, int hi)
    {
        // cout << s << endl;
        int last_op = -1;
        int depth = 0;
        for (int i = hi - 1; i >= lo; --i)
        {
            if (s[i] == ')')
                depth++;
            else if (s[i] == '(')
                depth--;

            if (depth == 0)
                if (s[i] == '*' || s[i] == '/')
                {
                    last_op = i;
                    break;
                }
        }

        if (last_op == -1)
            return primary_exp(s, lo, hi);
        else
        {
            if (s[last_op] == '*')
                return mul_exp(s, lo, last_op) * primary_exp(s, last_op + 1, hi);
            else
                return mul_exp(s, lo, last_op) / primary_exp(s, last_op + 1, hi);
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
        // cout << stoi(s.substr(lo, hi - lo)) << endl;
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