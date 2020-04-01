/*
 * @lc app=leetcode.cn id=127 lang=javascript
 *
 * [127] 单词接龙
 */

// @lc code=start
/**
 * @param {string} beginWord
 * @param {string} endWord
 * @param {string[]} wordList
 * @return {number}
 */
var ladderLength = function(beginWord, endWord, wordList) {
    const isSimilar = (a, b) => {
        let diff = 0;
        for(let i = 0; i < a.length; i++) {
            if (a.charAt(i) !== b.charAt(i)) diff ++;
            if (diff > 1) return false;
        }

        return true;
    }

    let index = wordList.indexOf(beginWord);

    if (index > -1) wordList.splice(index, 1);

    let queue = [beginWord];
    let res = 2;

    while(queue.length) {
        let size = queue.length;

        while(size --) {
            const front = queue.shift();
            for(let i = 0; i < wordList.length; i++) {
                if (!isSimilar(front, wordList[i])) continue;

                if (endWord == wordList[i]){
                    return res;
                }else {
                    queue.push(wordList[i]);
                }

                wordList.splice(i, 1);
                i --;
            }
        }

        res ++;
    }

    return 0;
};
// @lc code=end

