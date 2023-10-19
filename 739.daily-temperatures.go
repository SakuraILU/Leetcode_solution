/*
 * @lc app=leetcode id=739 lang=golang
 *
 * [739] Daily Temperatures
 */

// @lc code=start
type Temperature struct {
	day         int
	temperature int
}

func dailyTemperatures(temperatures []int) (nextwarmer []int) {
	nextwarmer = make([]int, len(temperatures))

	stack := make([]Temperature, 0)

	for i := len(temperatures) - 1; i >= 0; i-- {
		curtemp := temperatures[i]
		for len(stack) > 0 && curtemp >= stack[len(stack)-1].temperature {
			stack = stack[0 : len(stack)-1]
		}

		if len(stack) > 0 {
			nextwarmer[i] = stack[len(stack)-1].day - i
		} else {
			nextwarmer[i] = 0
		}

		stack = append(stack, Temperature{day: i, temperature: curtemp})
	}

	return
}

// @lc code=end
