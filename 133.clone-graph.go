/*
 * @lc app=leetcode id=133 lang=golang
 *
 * [133] Clone Graph
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

var visited map[*Node]*Node

func cloneGraph(node *Node) *Node {
	visited = make(map[*Node]*Node, 0)

	return dfs(node)
}

func dfs(node *Node) *Node {
	if node == nil {
		return nil
	} else if cnode, ok := visited[node]; ok {
		return cnode
	}

	cnode := &Node{
		Val:       node.Val,
		Neighbors: make([]*Node, 0),
	}
	visited[node] = cnode
	for _, neighbor := range node.Neighbors {
		cneighbor := dfs(neighbor)
		cnode.Neighbors = append(cnode.Neighbors, cneighbor)
	}

	return cnode
}

// @lc code=end
