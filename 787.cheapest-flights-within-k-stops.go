/*
 * @lc app=leetcode id=787 lang=golang
 *
 * [787] Cheapest Flights Within K Stops
 */

// @lc code=start
type State struct {
	node, leftstep int
}

var memo map[State]int

var graph [][]int
var weight [][]int
var _src, _dst int

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	graph, weight = buildGraph(n, flights)
	_src, _dst = src, dst

	memo = make(map[State]int)
	return dp(src, k+1)
}

func buildGraph(n int, flights [][]int) (graph [][]int, weight [][]int) {
	graph = make([][]int, n)
	weight = make([][]int, n)
	for i := 0; i < n; i++ {
		weight[i] = make([]int, n)
	}

	for _, flight := range flights {
		graph[flight[0]] = append(graph[flight[0]], flight[1])
		weight[flight[0]][flight[1]] = flight[2]
	}

	return
}

func dp(node int, leftstep int) (mval int) {
	if leftstep < 0 {
		return -1
	}
	if node == _dst {
		return 0
	}

	state := State{node: node, leftstep: leftstep}
	if v, ok := memo[state]; ok {
		return v
	}

	mval = math.MaxInt
	for _, neighbor := range graph[node] {
		val := dp(neighbor, leftstep-1)
		if val < 0 {
			continue
		}
		mval = min(mval, weight[node][neighbor]+val)
	}

	if mval == math.MaxInt {
		mval = -1
	}

	memo[state] = mval
	return mval
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// @lc code=end
