/**
 * @param {number[]} prices
 * @return {number}
 */
/**
 * 一次遍历
 */
// var maxProfit = function(prices) {
//     let res = 0
//     let inPrices = Infinity;
//     let outPrices;
//     for(let i=0;i<prices.length;i++) {
//         if(prices[i] < inPrices) {
//             inPrices = prices[i]
//             outPrices = -Infinity
//         }else if(outPrices < prices[i]){
//             outPrices = prices[i]
//             res = Math.max(res, outPrices - inPrices)
//         }
//     }
//     return res
// };
/**
 * 动态规划
 */
var maxProfit = function(prices) {
    let dp = [0];
    let min = prices[0];
    for(let i=1;i<prices.length;i++) {
        min = Math.min(min, prices[i])
        dp[i] = Math.max(dp[i-1], prices[i]- min)
    }
    return dp[dp.length-1]
};