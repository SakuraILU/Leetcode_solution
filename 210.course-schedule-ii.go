/*
 * @lc app=leetcode id=210 lang=golang
 *
 * [210] Course Schedule II
 */

// @lc code=start
var graph [][]int
var visited []bool
var onTrack []bool
var hasCircle bool

var res []int

func findOrder(numCourses int, prerequisites [][]int) []int {
	graph = buildGraph(numCourses, prerequisites)
	visited = make([]bool, numCourses)
	onTrack = make([]bool, numCourses)
	hasCircle = false

	res = make([]int, 0)

	for i := 0; i < numCourses; i++ {
		dfs(i)

		if hasCircle {
			return nil
		}
	}

	return res
}

func buildGraph(numCourses int, prerequisites [][]int) (graph [][]int) {
	graph = make([][]int, numCourses)

	for _, prerequisite := range prerequisites {
		graph[prerequisite[0]] = append(graph[prerequisite[0]], prerequisite[1])
	}

	return
}

func dfs(node int) {
	if hasCircle {
		return
	}

	if visited[node] {
		if onTrack[node] {
			hasCircle = true
		}
		return
	}

	visited[node] = true
	onTrack[node] = true

	for _, neighbor := range graph[node] {
		dfs(neighbor)
	}

	res = append(res, node)
	onTrack[node] = false
}

// @lc code=end
