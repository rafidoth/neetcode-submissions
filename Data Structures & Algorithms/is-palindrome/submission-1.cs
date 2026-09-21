public class Solution {
    public bool IsPalindrome(string s) {
        string s1 = string.Concat(s.Where(c=> char.IsLetterOrDigit(c))).ToLower();
        string s2 = string.Concat(s1.Reverse());
        if(s1 == s2) return true;
        return false;
    }
}
