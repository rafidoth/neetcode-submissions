func groupAnagrams(strs []string) [][]string {
		answer := make(map[string][]string)

	for _, str := range strs {
		bytes := []byte(str)
		comp := func(i, j int) bool {
			return bytes[i] < bytes[j]
		}
		sort.Slice(bytes, comp)
		sorted := string(bytes)
		answer[sorted] = append(answer[sorted], str)
	}
	var anagramGroups [][]string
	for _, anagramGroup := range answer {
		anagramGroups = append(anagramGroups, anagramGroup)
	}
	return anagramGroups

}
