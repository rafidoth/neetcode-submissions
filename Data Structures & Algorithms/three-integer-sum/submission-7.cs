public class Solution {
    public List<List<int>> ThreeSum(int[] nums) {
        Array.Sort(nums);
        List<List<int>> results = [];
        int n = nums.Length;

        for(int i=0; i<n; i++){
            if(i>0 && nums[i] == nums[i-1]) continue;
            int p1 = i+1;
            int p2 = n-1;
            while(p1 < p2){
                // ignore dups for both pointers
                int sum = nums[i] + nums[p1] + nums[p2];
                if(sum == 0){
                    results.Add([nums[i], nums[p1], nums[p2]]);
                }

                if(sum > 0) p2--;
                else p1++;

                while(p1 > i+1 && p1 < p2 && nums[p1] == nums[p1-1]) p1++; 
                while(p2 < n-1 && p1 < p2 && nums[p2] == nums[p2+1]) p2--; 
            }
        }

        return results;
    }
}
