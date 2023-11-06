/*
 * @lc app=leetcode id=1114 lang=cpp
 *
 * [1114] Print in Order
 */

// @lc code=start
class Foo
{
private:
    mutex lks[2];
    condition_variable cvs[2];

public:
    Foo()
    {
    }

    void first(function<void()> printFirst)
    {
        // printFirst() outputs "first". Do not change or remove this line.
        printFirst();
        cvs[0].notify_all();
    }

    void second(function<void()> printSecond)
    {
        unique_lock<mutex> lk(lks[0]);
        cvs[0].wait(lk);
        // printSecond() outputs "second". Do not change or remove this line.
        printSecond();
        cvs[1].notify_all();
    }

    void third(function<void()> printThird)
    {
        unique_lock<mutex> lk(lks[1]);
        cvs[1].wait(lk);
        // printThird() outputs "third". Do not change or remove this line.
        printThird();
    }
};
// @lc code=end
