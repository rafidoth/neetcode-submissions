class Solution {
    /**
     * @param {number[]} nums
     * @return {number}
     */
    longestConsecutive(nums) {
		let existanceBook = {}
		nums.forEach((val)=>{
			existanceBook[val] = true
		}) 
		let res = 0
		for(let i =0; i<nums.length; i++){
			let curr = nums[i]
			let cnt = 0
			while(existanceBook[curr]){
				cnt++
				curr++
			}
			res = Math.max(cnt, res)
		}
		return res
	}
}
