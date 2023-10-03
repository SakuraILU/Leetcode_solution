/*
 * @lc app=leetcode id=134 lang=golang
 *
 * [134] Gas Station
 */
// @lc code=start
func canCompleteCircuit(gas []int, cost []int) int {
	ngas := 0
	mingas, station := 0, 0
	for i := 0; i < len(gas); i++ {
		ngas += gas[i] - cost[i]
		if ngas < mingas {
			mingas = ngas
			station = i + 1
		}
	}

	if ngas < 0 {
		return -1
	}

	return station
}

// @lc code=end
