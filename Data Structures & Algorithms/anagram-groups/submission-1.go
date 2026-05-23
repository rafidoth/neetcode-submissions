func sortString (s string) string{
    r := []rune(s)
    sort.Slice(r, func(i,j int)bool{
        return r[i] < r[j]
    })
    return string(r)
}

func groupAnagrams(strs []string) [][]string {
    grps := make(map[string][]string)
    for _, x := range strs {
        sorted := sortString(x)
        if _, ok := grps[sorted]; ok {
            grps[sorted] = append(grps[sorted], x)
            continue
        }
        grps[sorted] = []string{x}
    }
    fmt.Println(grps, len(grps))
    result := make([][]string,0, len(grps))
    for _ , val := range grps {
        result = append(result , val)
    }
    return result
}
