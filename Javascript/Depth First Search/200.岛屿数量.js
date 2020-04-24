/*
 * @lc app=leetcode.cn id=200 lang=javascript
 *
 * [200] 岛屿数量
 *
 * https://leetcode-cn.com/problems/number-of-islands/description/
 *
 * algorithms
 * Medium (47.82%)
 * Likes:    539
 * Dislikes: 0
 * Total Accepted:    97.7K
 * Total Submissions: 198.6K
 * Testcase Example:  '[["1","1","1","1","0"],["1","1","0","1","0"],["1","1","0","0","0"],["0","0","0","0","0"]]'
 *
 * 给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
 * 
 * 岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
 * 
 * 此外，你可以假设该网格的四条边均被水包围。
 * 
 * 示例 1:
 * 
 * 输入:
 * 11110
 * 11010
 * 11000
 * 00000
 * 输出: 1
 * 
 * 
 * 示例 2:
 * 
 * 输入:
 * 11000
 * 11000
 * 00100
 * 00011
 * 输出: 3
 * 解释: 每座岛屿只能由水平和/或竖直方向上相邻的陆地连接而成。
 * 
 * 
 */

// @lc code=start
/**
 * @param {character[][]} grid
 * @return {number}
 */
var numIslands = function(grid) {
    let num = 0;
    if (grid && grid.length) {
        let x = grid.length - 1;
        let y = grid[0].length - 1;
        for(let i  = 0; i < grid.length; i++) {
            for(let j = 0; j < grid[i].length; j ++) {
                if (grid[i][j] === '1') {
                    num ++;
                    overturn(i, j);
                }
            }
        }
    
        function overturn(i, j) {
            if(i < 0 || j < 0 || i > x || j > y) return;
            if(grid[i][j] === '1') {
                grid[i][j] = '0'
                overturn(i, j-1)
                overturn(i-1, j)
                overturn(i+1, j)
                overturn(i, j+1)
            }
        }
    }

    return num;
};
// @lc code=end

