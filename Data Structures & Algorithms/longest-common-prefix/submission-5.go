func longestCommonPrefix(strs []string) string {
	prefMatched := strs[0]
	for _, str := range strs {
		prefMatched = prefMatched[:min(len(prefMatched), len(str))]
		for idx, char := range prefMatched {
			if char != rune(str[idx]) {
				prefMatched = prefMatched[:idx]
				break
			}
		}
	}
	return prefMatched



}
