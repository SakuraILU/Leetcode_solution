/*
 * @lc app=leetcode id=726 lang=golang
 *
 * [726] Number of Atoms
 */

// @lc code=start
func countOfAtoms(formula string) (res string) {
	cnts := dfs(formula)

	// sort keys
	keys := make([]string, 0)
	for k, _ := range cnts {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, key := range keys {
		res += string(key)
		if cnts[key] > 1 {
			res += strconv.Itoa(cnts[key])
		}
	}

	return res
}

func dfs(formula string) (cnts map[string]int) {
	if len(formula) < 2 {
		cnts = make(map[string]int)
		cnts[formula] = 1
		return
	} else if len(formula) == 2 {
		if unicode.IsUpper(rune(formula[0])) && unicode.IsLower(rune(formula[1])) {
			cnts = make(map[string]int)
			cnts[formula] = 1
			return
		}
	}

	i := len(formula) - 1
	for ; i >= 0 && unicode.IsDigit(rune(formula[i])); i-- {
	}

	k := 1
	if i < len(formula)-1 {
		k, _ = strconv.Atoi(formula[i+1 : len(formula)])
	}

	if formula[i] != ')' {
		j := i
		if unicode.IsLower(rune(formula[i])) {
			j--
		}
		if j > 0 {
			cnts = dfs(formula[0:j])
		} else {
			cnts = make(map[string]int)
		}

		cnts[formula[j:i+1]] += k
	} else {
		depth := 0
		j := i
		for ; j >= 0; j-- {
			if formula[j] == ')' {
				depth--
			} else if formula[j] == '(' {
				depth++

				if depth == 0 {
					break
				}
			}
		}

		if j > 0 {
			cnts = dfs(formula[0:j])
		} else {
			cnts = make(map[string]int)
		}

		substr_cnts := dfs(formula[j+1 : i])
		for key, val := range substr_cnts {
			cnts[key] += val * k
		}
	}

	return
}

// @lc code=end
