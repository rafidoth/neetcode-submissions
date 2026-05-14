func majorityElement(nums []int) int {
    counts := make(map[int]int)

	for _, num := range nums {
		counts[num]++
	}

	n := len(nums) / 2
	for key, val := range counts {
		if val > n {
			return key
		}
	}
	return -1

}
