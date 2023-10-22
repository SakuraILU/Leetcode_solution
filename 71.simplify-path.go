/*
 * @lc app=leetcode id=71 lang=golang
 *
 * [71] Simplify Path
 */

// @lc code=start
import (
	"strings"
)

const SEP = "/"

func simplifyPath(path string) string {
	stack := make([]string, 0)

	names := strings.Split(path, SEP)

	for _, name := range names {
		if name == "" || name == "." {
			continue
		}

		if name == ".." {
			if len(stack) > 0 {
				stack = stack[0 : len(stack)-1]
			}
		} else {
			stack = append(stack, name)
		}
	}

	return SEP + strings.Join(stack, SEP)
}

// @lc code=end
