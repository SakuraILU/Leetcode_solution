/*
 * @lc app=leetcode id=207 lang=golang
 *
 * [207] Course Schedule
 */

// @lc code=start
var graph [][]int
var visited []bool
var onTrack []bool
var hasCircle bool

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph = buildGraph(numCourses, prerequisites)
	visited = make([]bool, numCourses)
	onTrack = make([]bool, numCourses)
	hasCircle = false

	for i := 0; i < numCourses; i++ {
		dfs(i)
		if hasCircle {
			return false
		}
	}

	return true
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

	onTrack[node] = false
}

// @lc code=end
