func productExceptSelf(nums []int) []int {

	n := len(nums)
	pref := make([]int, n)
	suff := make([]int, n)

	pref[0], suff[n-1] = 1,1
	
	for i:=1; i<n; i++ {
		pref[i] = pref[i-1]*nums[i-1]
	}

	for i:= n-2; i>=0; i-- {
		suff[i] = suff[i+1]*nums[i+1]
	}

	res := make([]int, n)
	for i:= 0; i<n; i++ {
		res[i] = pref[i] * suff[i]
	}

	return res
	
}
