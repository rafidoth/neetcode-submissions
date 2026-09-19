public class Solution {
    public int CountComponents(int n, int[][] edges) {
        Dictionary<int, List<int>> adjMat = new();
        for(int i =0 ;i<n; i++){
            adjMat[i] = new List<int>();
        }

        foreach(int[] edge in edges){
            int u = edge[0];
            int v = edge[1];
            adjMat[u].Add(v); 
            adjMat[v].Add(u);
        }

        var visited = new HashSet<int>();
        int count = 0;
        for (int i = 0; i < n; i++){
            if(!visited.Contains(i)){
                Dfs(i, adjMat, visited);
                count++;
            }
        }
        return count;
    }

    private void Dfs(int start, Dictionary<int,List<int>> adjMat, HashSet<int> visited){
        if(visited.Contains(start))
            return;
        visited.Add(start); 
        foreach(int neigh in adjMat[start]){
            Dfs(neigh, adjMat, visited);
        }
        return;
    }
}
