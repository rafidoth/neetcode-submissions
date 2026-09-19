public class Solution {
    private HashSet<int> visited = new();
    private Dictionary<int,List<int>> matrix = new();

    public bool ValidTree(int n, int[][] edges) {
        if (edges.Length != n - 1) return false;

        Dictionary<int, List<int>> adj = new();
        for(int i =0; i<n; i++)
            adj[i] = new List<int>();
        
        foreach(int[] edge in edges){
            int u = edge[0];
            int v = edge[1];
            adj[u].Add(v);
            adj[v].Add(u);
        }
        matrix = adj;

        if(!Dfs(0, -1)) return false;
        return visited.Count == n;
    }

    private bool Dfs(int start, int parent){
        if(visited.Contains(start)) return false;

        visited.Add(start);
        foreach(int neigh in matrix[start]){
            if(neigh != parent){
                if(!Dfs(neigh, start)){
                    return false;
                }
            }
        }
        return true;
    }
}