package ArrayList

import (
	"strings"
)

// 回溯法，每一层是一个新的数组(下一个数字对应的字母数组)
var letters []string
var let	string
var letterMap map[string][]string
var digitsList []string

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	letters = []string{}
	let = ""
	digitsList = strings.Split(digits, "")
	letterMap = make(map[string][]string)
	letterMap["2"] = []string{"a", "b", "c"}
	letterMap["3"] = []string{"d", "e", "f"}
	letterMap["4"] = []string{"g", "h", "i"}
	letterMap["5"] = []string{"j", "k", "l"}
	letterMap["6"] = []string{"m", "n", "o"}
	letterMap["7"] = []string{"p", "q", "r", "s"}
	letterMap["8"] = []string{"t", "u", "v"}
	letterMap["9"] = []string{"w", "x", "y", "z"}
	backTrackLetters(letterMap[digitsList[0]], digitsList)
	return letters
}

func backTrackLetters(nums, digits []string)  {
	if len(let) == len(digits) {
		letters = append(letters, let)
		return
	}
	for i:=0; i<len(nums); i++ {
		let += nums[i]
		if len(let) < len(digits) {
			backTrackLetters(letterMap[digits[len(let)]], digits)
		} else {
			backTrackLetters([]string{}, digits)
		}
		let = let[:len(let)-1]
	}
}
