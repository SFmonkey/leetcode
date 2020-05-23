/**
 * @param {number[]} digits
 * @return {number[]}
 */
/**
 * 遍历数组给最后一个元素加一
 * 若果最后一个数组为9 则需要给前一个元素加一
 * 如果最后一个元素不为9 直接给最后一个元素加一
 * 如果第一个元素为9，则需要给数组头添加1
 */
var plusOne = function(digits) {
    for(let i=digits.length-1;i>= 0;i--) {
        if(digits[i] == 9) {
            digits[i] = 0 
            if(i == 0) {
                digits.splice(0, 0, 1)
                return digits
            }
        }
        else{
            digits[i]++
            return digits
        }
    }
    
};