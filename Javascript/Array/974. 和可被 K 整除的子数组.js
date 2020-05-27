/**
 * @param {number[]} A
 * @param {number} K
 * @return {number}
 */
/**
 * 前缀和思想
 * 使用preArr 存储前N项和对K的余数  key为前N项的和对K的余数  value为Key出现的次数
 * sum记录前N项和对K的余数，当sum小于0时需要将其加 K
 * 如果之前出现过和当前preArr[sum]相同的值，则从该项到第i项的和可以整除K，所以count+= preArr[sum]
 * 因为preArr[sum]又出现一次，所以preArr[sum]++
 * 
 * 如果preArr[sum]没有出现过 将其置为preArr[sum] = 1
 */
var subarraysDivByK = function(A, K) {
    let preArr = {0: 1}
    let sum = 0
    let count = 0
    for(let i=0;i<A.length;i++) {
        sum = (sum + A[i]) % K
        if(sum < 0) {
            sum+= K
        }
        if(preArr[sum]) {
            count+= preArr[sum]
            preArr[sum]++
        }else{
            preArr[sum] = 1
        }
    }
    return count
};