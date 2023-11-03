/*
 * @lc app=leetcode id=127 lang=golang
 *
 * [127] Word Ladder
 */

// @lc code=start

var _wordlist []string

func ladderLength(beginWord string, endWord string, wordList []string) int {
	_wordlist = wordList

	queue := make([]string, 0)
	nqueue := make([]string, 0)
	visited := make(map[string]bool)

	queue = append(queue, beginWord)
	length := 0
	for len(queue) > 0 {
		length++
		for _, node := range queue {
			if node == endWord {
				return length
			}

			if visited[node] {
				continue
			}
			visited[node] = true

			neighbors := getNeighbors(node)
			for _, neighbor := range neighbors {
				nqueue = append(nqueue, neighbor)
			}
		}
		queue = nqueue
		nqueue = make([]string, 0)
	}

	return 0
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
