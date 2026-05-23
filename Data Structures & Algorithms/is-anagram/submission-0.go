func sortString(s string) string{
    runes := []rune(s)
    sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
    s = string(runes)
    return s
}   


func isAnagram(s string, t string) bool {
    s = sortString(s)
    t = sortString(t)
    return s==t
}
