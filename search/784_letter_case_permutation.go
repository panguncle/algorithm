package search

import "strings"

// LetterCasePermutation: leetcode.784
func LetterCasePermutation(s string) []string {
	var ret []string
	if len(s) == 0 {
		return ret
	}

	var cur string
	letterCasePermutation(&ret, cur, s, 0)
	return ret
}

func letterCasePermutation(ret *[]string, cur, s string, level int) {
	if len(cur) > len(s) {
		return
	}

	if len(cur) == len(s) {
		*ret = append(*ret, cur)
		return
	}

	for i := level; i < len(s); i++ {
		v := string(s[i])
		if s[i] >= 'A' && s[i] <= 'z' {
			cur += strings.ToLower(v)
			letterCasePermutation(ret, cur, s, i+1)
			cur = cur[0 : len(cur)-1]
			cur += strings.ToUpper(v)
			letterCasePermutation(ret, cur, s, i+1)
			cur = cur[0 : len(cur)-1]
		} else {
			cur += v
			letterCasePermutation(ret, cur, s, i+1)
			cur = cur[0 : len(cur)-1]
		}
	}
}
