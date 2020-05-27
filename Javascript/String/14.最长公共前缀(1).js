/**
 * @param {string[]} strs
 * @return {string}
 */
var longestCommonPrefix = function(strs) {
    let i= 0;
    let res = ""
    while(strs[0]){
        let str = strs[0].slice(0, i+1)
        for(let j=0;j<strs.length;j++) {
            if(strs[j].indexOf(str) !== 0) {
                return res
            }
        }
        if(i == strs[0].length) {
            return strs[0]
        }
        res = str
        i++
    }
    return res
};