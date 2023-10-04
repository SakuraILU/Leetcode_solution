/*
 * @lc app=leetcode id=427 lang=golang
 *
 * [427] Construct Quad Tree
 */

// @lc code=start
/**
 * Definition for a QuadTree node.
 * type Node struct {
 *     Val bool
 *     IsLeaf bool
 *     TopLeft *Node
 *     TopRight *Node
 *     BottomLeft *Node
 *     BottomRight *Node
 * }
 */

var _grid [][]int

func construct(grid [][]int) *Node {
	_grid = grid

	return buildGraph(0, 0, len(grid)-1, len(grid[0])-1)
}

func buildGraph(x1, y1, x2, y2 int) (root *Node) {
	if x1 == x2 && y1 == y2 {
		return &Node{
			Val:    _grid[x1][y1] == 1,
			IsLeaf: true,
		}
	}

	x, y := x1+(x2-x1)/2, y1+(y2-y1)/2
	tltree := buildGraph(x1, y1, x, y)
	trtree := buildGraph(x1, y+1, x, y2)
	bltree := buildGraph(x+1, y1, x2, y)
	brtree := buildGraph(x+1, y+1, x2, y2)

	if tltree.IsLeaf && trtree.IsLeaf && bltree.IsLeaf && brtree.IsLeaf {
		if (tltree.Val && trtree.Val && bltree.Val && brtree.Val) || (!tltree.Val && !trtree.Val && !bltree.Val && !brtree.Val) {
			root = &Node{
				Val:    tltree.Val,
				IsLeaf: true,
			}
			return
		}
	}

	root = &Node{
		IsLeaf:      false,
		TopLeft:     tltree,
		TopRight:    trtree,
		BottomLeft:  bltree,
		BottomRight: brtree,
	}

	return
}

// @lc code=end
