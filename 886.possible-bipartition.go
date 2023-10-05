/*
 * @lc app=leetcode id=886 lang=golang
 *
 * [886] Possible Bipartition
 */

// @lc code=start
var graph [][]int
var visited []bool
var color []bool
var isbipart bool

func possibleBipartition(n int, dislikes [][]int) bool {
	graph = buildGraph(n, dislikes)
	visited = make([]bool, n+1)
	color = make([]bool, n+1)
	isbipart = true

	for i := 1; i <= n; i++ {
		dfs(i)
		if !isbipart {
			return false
		}
	}

	return true
}

func buildGraph(n int, dislikes [][]int) (graph [][]int) {
	graph = make([][]int, n+1)
	for _, dislike := range dislikes {
		graph[dislike[0]] = append(graph[dislike[0]], dislike[1])
		graph[dislike[1]] = append(graph[dislike[1]], dislike[0])
	}

	return
}

func dfs(node int) {
	if !isbipart {
		return
	}
	if visited[node] {
		return
	}
	visited[node] = true

	for _, neighbor := range graph[node] {
		if !visited[neighbor] {
			color[neighbor] = !color[node]
			dfs(neighbor)
		} else {
			if color[neighbor] == color[node] {
				isbipart = false
			}
		}
	}
}

// @lc code=end
