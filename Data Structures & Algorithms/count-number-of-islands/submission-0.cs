public class Solution {
    private const char water = '0';
    private const char land = '1';
    private HashSet<(int,int)> visited = new();
    public int NumIslands(char[][] grid) {
        int rowsCount = grid.Length;
        int colsCount = grid[0].Length;
        int count = 0;
        for(int i =0; i<rowsCount; i++){
            for(int j=0; j<colsCount; j++){
                if(grid[i][j]==land && !visited.Contains((i,j))){
                    findIsland(i,j, grid);
                    count++;
                }
            }
        }

        return count;

    }

    private void findIsland(int row, int col, char[][] grid){
        if(visited.Contains((row,col))){
            return;
        }
        
        if(grid[row][col] == water){
            return;
        }

        int rowsCount = grid.Length;
        int colsCount = grid[0].Length;

        visited.Add((row,col));

        if(row - 1 >= 0){
            findIsland(row-1, col, grid);
        }

        if(row +1 < rowsCount){
            findIsland(row+1, col, grid);
        }

        if(col-1 >= 0){
            findIsland(row, col -1, grid);
        }

        if(col + 1 < colsCount){
            findIsland(row, col+1, grid);
        }

    }
}
