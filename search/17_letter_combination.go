package search

var phoneNumMapper = map[string][]string{
	"2": {"a", "b", "c"},
	"3": {"d", "e", "f"},
	"4": {"g", "h", "i"},
	"5": {"j", "k", "l"},
	"6": {"m", "n", "o"},
	"7": {"p", "q", "r", "s"},
	"8": {"t", "u", "v"},
	"9": {"w", "x", "y", "z"},
}

// LetterCombinations: leetcode.17
func LetterCombinations(digits string) []string {

	var ret []string
	if len(digits) == 0 {
		return ret
	}

	var cur string
	letterCombinations(&ret, cur, digits, 0)
	return ret
}

func letterCombinations(ret *[]string, cur, digits string, level int) {
	if len(cur) == len(digits) {
		*ret = append(*ret, cur)
		return
	}

	for i := level; i < len(digits); i++ {
		v := string(digits[i])
		mapper := phoneNumMapper[v]
		for _, c := range mapper {
			cur += c
			letterCombinations(ret, cur, digits, i+1)
			cur = cur[:len(cur)-1]
		}
	}
}
