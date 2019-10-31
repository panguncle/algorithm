package search

// GenrateParentheses: leetcode.22
func GenerateParentheses(n int) []string {
	var ret []string
	if n == 0 {
		return ret
	}
	var cur string
	generateParentheses(&ret, cur, n, n)
	return ret
}

func generateParentheses(ret *[]string, cur string, leftLeft, rightLeft int) {
	if leftLeft < 0 || rightLeft < 0 || rightLeft < leftLeft {
		return
	}

	if leftLeft == 0 && rightLeft == 0 {
		*ret = append(*ret, cur)
		return
	}

	if leftLeft > 0 {
		generateParentheses(ret, cur+"(", leftLeft-1, rightLeft)
	}

	// 不使用rightLeft
	if rightLeft > leftLeft {
		generateParentheses(ret, cur+")", leftLeft, rightLeft-1)
	}
}
