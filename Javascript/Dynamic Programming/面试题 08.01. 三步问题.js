/**
 * @param {number} n
 * @return {number}
 */
var waysToStep = function(n) {
    let dp = []
    dp[0] = 0
    dp[1] = 1
    dp[2] = 2
    dp[3] = 4
    for(let i=4;i<=n;i++) {
        dp[i] = (dp[i-3]+ dp[i-2]+ dp[i-1]) % 1000000007
    }
    return dp[n]
};