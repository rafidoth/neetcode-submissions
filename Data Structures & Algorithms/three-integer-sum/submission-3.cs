public class Solution {
    public List<List<int>> ThreeSum(int[] nums) {
        Array.Sort(nums);
        List<List<int>> results = [];
        int n = nums.Length;
        Dictionary<int,int> map = new();
        foreach(int num in nums){
            if(!map.ContainsKey(num)){
                map[num] = 0;
            }
            map[num]++;
        }

        for(int i = 0; i<n; i++){
            map[nums[i]]--;
            if(i > 0 && nums[i] == nums[i-1]) continue;
            for(int j = i+1; j<n; j++){
                map[nums[j]]--;
                if(j > i+1 && nums[j] == nums[j-1]) continue;

                int sum = nums[i] + nums[j];
                int num3 = -sum; 
                if(map.ContainsKey(num3) && map[num3] > 0 ){
                    results.Add([nums[i], nums[j], num3]);
                }
            }
            for (int j = i + 1; j < nums.Length; j++) {
                map[nums[j]]++;
            }
        }
        return results;
    }
}
