func twoSum(nums []int, target int) []int {
    need := make(map[int]int)
    for i , x := range nums{
        if _, ok := need[x]; ok {
            return []int{need[x], i}
        }
        need[target-x] = i
    }
    return []int{-1,-1}
}
