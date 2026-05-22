class Solution {
    /**
     * @param {number[]} nums
     * @return {number}
     */
    longestConsecutive(nums) {
		if(nums.length === 0){
			return 0 
		}

		if(nums.length === 1 ){
			return 1
		}
		
		nums.sort((a,b)=> a-b)

		let res = 0
		let i =0
		let cnt = 0
		while (i + 1  < nums.length){
			while(nums[i+1]=== nums[i]+1){
				cnt++
				i++
			}
			res = Math.max(res, cnt+1)
			if(nums[i+1] - nums[i] > 1){
				cnt = 0
			}
			i++
		}
		return res
	}
}
