func majorityElement(nums []int) int {
   	var bitCounts [32]int

	for _, num := range nums {
		for i := range 32 {
			ithBit := (num >> i) & 1
			bitCounts[i] += ithBit
		}
	}

	half := len(nums) / 2
	var result int32
	for i, count := range bitCounts {
		if count > half {
			result |= (1 << i)
		}
	}
	return int(result)


}
