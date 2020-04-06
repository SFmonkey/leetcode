/*
 * @lc app=leetcode.cn id=54 lang=javascript
 *
 * [54] 螺旋矩阵
 *
 * https://leetcode-cn.com/problems/spiral-matrix/description/
 *
 * algorithms
 * Medium (39.24%)
 * Likes:    334
 * Dislikes: 0
 * Total Accepted:    50.4K
 * Total Submissions: 128K
 * Testcase Example:  '[[1,2,3],[4,5,6],[7,8,9]]'
 *
 * 给定一个包含 m x n 个元素的矩阵（m 行, n 列），请按照顺时针螺旋顺序，返回矩阵中的所有元素。
 * 
 * 示例 1:
 * 
 * 输入:
 * [
 * ⁠[ 1, 2, 3 ],
 * ⁠[ 4, 5, 6 ],
 * ⁠[ 7, 8, 9 ]
 * ]
 * 输出: [1,2,3,6,9,8,7,4,5]
 * 
 * 
 * 示例 2:
 * 
 * 输入:
 * [
 * ⁠ [1, 2, 3, 4],
 * ⁠ [5, 6, 7, 8],
 * ⁠ [9,10,11,12]
 * ]
 * 输出: [1,2,3,4,8,12,11,10,9,5,6,7]
 * 
 * 
 */

// @lc code=start
/**
 * @param {number[][]} matrix
 * @return {number[]}
 */
var spiralOrder = function(matrix) {
    if (matrix.length < 1) return [];
    let res = [];
    
    if (matrix[0].length == 1) {
        for(let i = 0; i < matrix.length; i ++) {
            res.push(matrix[i][0]);
        }

        return res;
    }
    
    let m = matrix[0].length;
    let n = matrix.length;
    let left = 0,
        right = m - 1,
        up = 0,
        down = n - 1;

    let x = 0;
    let y = 0;
    let state = "right";
    up ++;
    for (var i = 0; i < m * n; i++) {
        res[i] = matrix[x][y];
        switch (state) {
            case "right":
                y ++;
                if (y == right) {
                    state = "down";
                    right --;
                }
                break;
            case "down":
                x++;
                if (x == down) {
                    state = "left";
                    down --;
                }
                break;
            case "left":
                y --;
                if (y == left) {
                    state = 'up';
                    left ++;
                }
                break;
            case "up":
                x --;
                if (x == up) {
                    state = 'right';
                    up ++;
                }
                break;
        }
    }

    return res;
};
// @lc code=end

