/*
 * @lc app=leetcode id=785 lang=golang
 *
 * [785] Is Graph Bipartite?
 */

// @lc code=start
var _graph [][]int
var visited []bool
var color []bool
var isbipart bool

func isBipartite(graph [][]int) bool {
	_graph = graph
	visited = make([]bool, len(graph))
	color = make([]bool, len(graph))
	isbipart = true

	for i := 0; i < len(graph); i++ {
		if !visited[i] {
			dfs(i)
			if !isbipart {
				return false
			}
		}
	}

	return true
}

func dfs(node int) {
	fmt.Println(node, color[node])
	if !isbipart {
		return
	}
	if visited[node] {
		return
	}
	visited[node] = true

	for _, neighbor := range _graph[node] {
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
