/*
 * @lc app=leetcode id=399 lang=golang
 *
 * [399] Evaluate Division
 */

// @lc code=start
type Adjence struct {
	Val  float64
	Node string
}

var graph map[string][]Adjence
var onTrack map[string]bool
var track float64

var find bool
var res []float64

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	graph = buildGraph(equations, values)
	onTrack = make(map[string]bool, len(graph))
	track = 1

	res = make([]float64, 0)

	for _, query := range queries {
		_, ok0 := graph[query[0]]
		_, ok1 := graph[query[1]]
		if !ok0 || !ok1 {
			res = append(res, -1.0)
			continue
		}

		find = false
		dfs(query[0], query[1])
		if !find {
			res = append(res, -1.0)
		}
	}

	return res
}

func buildGraph(equations [][]string, values []float64) (graph map[string][]Adjence) {
	graph = make(map[string][]Adjence)
	for idx, equation := range equations {
		graph[equation[0]] = append(graph[equation[0]], Adjence{Val: values[idx], Node: equation[1]})
		graph[equation[1]] = append(graph[equation[1]], Adjence{Val: 1 / values[idx], Node: equation[0]})
	}
	return
}

func dfs(node string, target string) {
	if find {
		return
	}
	if node == target {
		res = append(res, track)
		find = true
		return
	}

	if onTrack[node] {
		return
	}
	onTrack[node] = true

	for _, adj := range graph[node] {
		track = track * adj.Val
		dfs(adj.Node, target)
		track = track / adj.Val
	}

	onTrack[node] = false

	return
}

// @lc code=end
