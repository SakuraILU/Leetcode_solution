/*
 * @lc app=leetcode id=2097 lang=golang
 *
 * [2097] Valid Arrangement of Pairs
 */

// @lc code=start
type Edge struct {
	v0, v1 int
}

var graph map[int][]int
var visited map[Edge]bool
var start int
var result [][]int

func validArrangement(pairs [][]int) [][]int {
	graph, visited, start = buildGraph(pairs)
	result = make([][]int, 0)

	dfs(start)

	reverse(result)

	return result
}

func buildGraph(pairs [][]int) (map[int][]int, map[Edge]bool, int) {
	graph := make(map[int][]int)
	visited := make(map[Edge]bool)
	in := make(map[int]int)
	out := make(map[int]int)

	for _, pair := range pairs {
		graph[pair[0]] = append(graph[pair[0]], pair[1])
		visited[Edge{v0: pair[0], v1: pair[1]}] = false
		out[pair[0]]++
		in[pair[1]]++
	}

	start := pairs[0][0]
	for node, outdegree := range out {
		if outdegree == in[node]+1 {
			start = node
			break
		}
	}

	return graph, visited, start
}

func dfs(node int) {
	for _, neighbor := range graph[node] {
		edge := Edge{v0: node, v1: neighbor}
		if !visited[edge] {
			visited[edge] = true
			dfs(neighbor)
			result = append(result, []int{node, neighbor})
		}
	}
}

func reverse(arr [][]int) {
	left, right := 0, len(arr)-1
	for left <= right {
		tmp := arr[left]
		arr[left] = arr[right]
		arr[right] = tmp

		left++
		right--
	}
}

// @lc code=end
