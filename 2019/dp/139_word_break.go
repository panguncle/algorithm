package dp

/*
WordBreak :139
dp[i]: 前i个字符能否根据wordDict中的词被分词
dp[i] =
*/
func WordBreak(s string, wordDict []string) bool {

	_wordDict := make(map[string]bool, len(wordDict))
	for _, item := range wordDict {
		_wordDict[item] = true
	}

	mem := make(map[string]bool)
	return wordBreakHelper(s, _wordDict, mem)
}

func wordBreakHelper(s string, wordDict, mem map[string]bool) bool {
	if v, exists := mem[s]; exists {
		return v
	}

	if _, exists := wordDict[s]; exists {
		mem[s] = true
		return true
	}

	for j := 1; j < len(s); j++ {
		left := string(s[0:j])
		right := string(s[j:])
		if mem[right] && wordBreakHelper(left, wordDict, mem) {
			mem[s] = true
			return true
		}
	}
	mem[s] = false
	return false
}
