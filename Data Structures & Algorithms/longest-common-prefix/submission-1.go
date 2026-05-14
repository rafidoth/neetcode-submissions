func longestCommonPrefix(strs []string) string {
		prefMatched := strs[0]
	for ptr := range 200 {
		for _, str := range strs {
			if ptr == len(str) {
				return prefMatched[:ptr]
			}
			if str[ptr] != prefMatched[ptr] {
				return prefMatched[:ptr]

			}
		}
	}
	return prefMatched

}
