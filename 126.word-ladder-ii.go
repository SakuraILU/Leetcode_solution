/*
 * @lc app=leetcode id=126 lang=golang
 *
 * [126] Word Ladder II
 */

// @lc code=start
var _wordlist []string
var _beginWord string

var parents map[string][]string

var track []string
var res [][]string

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	_wordlist = wordList
	_beginWord = beginWord

	parents = make(map[string][]string)

	track = make([]string, 0)
	res = make([][]string, 0)

	queue := make([]string, 0)
	nqueue := make([]string, 0)
	visited := make(map[string]bool)

	queue = append(queue, beginWord)
	visited[beginWord] = true
	for len(queue) > 0 {
		for _, node := range queue {
			if node == endWord {
				dfs(endWord)
				return res
			}

			neighbors := getNeighbors(node)
			for _, neighbor := range neighbors {
				if visited[neighbor] {
					continue
				}
				visited[neighbor] = true

				nqueue = append(nqueue, neighbor)
			}
		}

		for _, child := range nqueue {
			for _, parent := range queue {
				if isNeighbor(child, parent) {
					parents[child] = append(parents[child], parent)
				}
			}
		}

		queue = nqueue
		nqueue = make([]string, 0)
	}

	return [][]string{}
}

func dfs(node string) {
	track = append(track, node)

	if node == _beginWord {
		res = append(res, reverse(track))
	} else {
		for _, neighbor := range parents[node] {
			dfs(neighbor)
		}
	}

	track = track[0 : len(track)-1]
}

func reverse(strs []string) (res []string) {
	res = make([]string, 0)
	for i := len(strs) - 1; i >= 0; i-- {
		res = append(res, strs[i])
	}

	return res
}

func getNeighbors(node string) []string {
	neighbors := make([]string, 0)

	for _, word := range _wordlist {
		if isNeighbor(node, word) {
			neighbors = append(neighbors, word)
		}
	}

	return neighbors
}

func isNeighbor(node1, node2 string) bool {
	if len(node1) != len(node2) {
		return false
	}

	diff := 0
	for i := 0; i < len(node1); i++ {
		if node1[i] != node2[i] {
			diff++
		}

		if diff > 1 {
			return false
		}
	}

	return true
}

// @lc code=end
