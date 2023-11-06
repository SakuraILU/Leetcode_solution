/*
 * @lc app=leetcode id=1116 lang=cpp
 *
 * [1116] Print Zero Even Odd
 */

// @lc code=start

#include <iostream>
#include <pthread.h>
using namespace std;

class ZeroEvenOdd
{
private:
    int n;
    int turn;
    pthread_cond_t cv;
    pthread_mutex_t mutex;

public:
    ZeroEvenOdd(int n)
    {
        this->n = n;
        this->turn = 0;
        pthread_cond_init(&cv, NULL);
        pthread_mutex_init(&mutex, NULL);
    }

    // printNumber(x) outputs "x", where x is an integer.
    void zero(function<void(int)> printNumber)
    {
        for (int i = 0; i < n; i++)
        {
            pthread_mutex_lock(&mutex);

            while (!(turn % 4 == 0 || turn % 4 == 2))
            {
                pthread_cond_wait(&cv, &mutex);
            }

            printNumber(0);
            turn++;

            pthread_cond_broadcast(&cv);
            pthread_mutex_unlock(&mutex);
        }
    }

    void even(function<void(int)> printNumber)
    {
        for (int i = 2; i <= n; i = i + 2)
        {
            pthread_mutex_lock(&mutex);

            while (!(turn % 4 == 3))
            {
                pthread_cond_wait(&cv, &mutex);
            }

            printNumber(i);
            turn++;

            pthread_cond_broadcast(&cv);
            pthread_mutex_unlock(&mutex);
        }
    }

    void odd(function<void(int)> printNumber)
    {
        for (int i = 1; i <= n; i = i + 2)
        {
            pthread_mutex_lock(&mutex);

            while (!(turn % 4 == 1))
            {
                pthread_cond_wait(&cv, &mutex);
            }

            printNumber(i);
            turn++;

            pthread_cond_broadcast(&cv);
            pthread_mutex_unlock(&mutex);
        }
    }
};
// @lc code=end
