func productExceptSelf(nums []int) []int {
	var result []int
	allProduct := 1
	zeroCnt := 0
	for _, x := range nums {
		if x== 0 {
			zeroCnt++
			continue
		}
		allProduct *= x
	}
	if zeroCnt > 1 {
		r := make([]int, len(nums))
		result = r
	}else if zeroCnt == 0 {
		var r []int
		for _, x := range nums {
			r = append(r, allProduct/x)
		}
		result = r
	}else{
		// zero count 1	
		r := make([]int, len(nums))
		allProductWithoutZero := 1
		zeroIdx := -1
		for idx , x := range nums{
			if x == 0 {
				zeroIdx = idx
				continue
			}
			allProductWithoutZero *= x
		}
		r[zeroIdx] = allProductWithoutZero
		result = r
	}

	
	return result

}
