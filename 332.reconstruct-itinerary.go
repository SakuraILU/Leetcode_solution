/*
 * @lc app=leetcode id=332 lang=golang
 *
 * [332] Reconstruct Itinerary
 */

// @lc code=start

import (
	"fmt"
	"sort"
)

type Edge struct {
	v0, v1 string
}

var graph map[string][]string
var edges map[Edge]int
var queue []string

func findItinerary(tickets [][]string) []string {
	graph, edges = buildGraph(tickets)
	queue = make([]string, 0)

	dfs("JFK")
	reverse(queue)

	return queue
}

func buildGraph(tickets [][]string) (graph map[string][]string, edges map[Edge]int) {
	graph = make(map[string][]string)
	edges = make(map[Edge]int)

	for _, ticket := range tickets {
		graph[ticket[0]] = append(graph[ticket[0]], ticket[1])
		edges[Edge{v0: ticket[0], v1: ticket[1]}]++
	}

	for key, adj := range graph {
		sort.Strings(adj)
		fmt.Println(key, adj)
	}

	return
}

func dfs(node string) {
	for _, neighbor := range graph[node] {
		edge := Edge{v0: node, v1: neighbor}
		if edges[edge] > 0 {
			edges[edge]--
			dfs(neighbor)
		}
	}

	queue = append(queue, node)
}

func reverse(arr []string) {
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
