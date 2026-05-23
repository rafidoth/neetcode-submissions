func hasDuplicate(nums []int) bool {
    f := make(map[int]bool)
    for _, x := range nums {
        if f[x] {
            return true
        }
        f[x] = true
    }
    return false
}
