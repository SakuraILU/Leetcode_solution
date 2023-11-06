/*
 * @lc app=leetcode id=1115 lang=cpp
 *
 * [1115] Print FooBar Alternately
 */

// @lc code=start
#include <mutex>
#include <condition_variable>
using namespace std;

class FooBar
{
private:
    int n;

    mutex mlk;
    condition_variable cv;
    int cnt;

public:
    FooBar(int n)
    {
        this->n = n;
        this->cnt = 0;
    }

    void foo(function<void()> printFoo)
    {
        unique_lock<mutex> lk(mlk);
        for (int i = 0; i < n; i++)
        {
            cv.wait(lk, [this]() -> bool
                    { return cnt % 2 == 0; });

            printFoo();
            // printf("foo\n");
            cnt++;
            // printFoo() outputs "foo". Do not change or remove this line.
            cv.notify_all();
        }
    }

    void bar(function<void()> printBar)
    {
        unique_lock<mutex> lk(mlk);
        for (int i = 0; i < n; i++)
        {
            cv.wait(lk, [this]()
                    { return cnt % 2 == 1; });
            // printBar() outputs "bar". Do not change or remove this line.
            printBar();
            // printf("bar\n");
            cnt++;

            cv.notify_all();
        }
    }
};
// @lc code=end
