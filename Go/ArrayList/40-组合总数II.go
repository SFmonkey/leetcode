package ArrayList

import (
	"fmt"
	"sort"
)

// 回溯法(i+1)，排序后去重
var comb [][]int
var cand []int
var candTmp []int
var tt int

func combinationSum2(candidates []int, target int) [][]int {
	//[1,1,6],[1,2,5],[1,7],[2,6]]
	tt = target
	comb = [][]int{}
	cand = []int{}
	sort.Ints(candidates)
	btCombination(candidates, 0)
	fmt.Println(comb)
	return comb
}

func btCombination(nums []int, startIndex int)  {
	if sumInt() == tt {
		candTmp = make([]int, len(cand))			// todo: []int{} 不行？
		copy(candTmp, cand)
		comb = append(comb, candTmp)
		return
	}
	for i:=startIndex; i<len(nums); i++ {
		// todo: 理解去重
		if i>startIndex && nums[i] == nums[i-1] {
			continue
		}
		cand = append(cand, nums[i])
		if sumInt() > tt {
			cand = cand[:len(cand)-1]
			continue
		}
		btCombination(nums, i+1)
		cand = cand[:len(cand)-1]
	}
}

func sumInt() (total int) {
	for i:=0; i<len(cand); i++ {
		total += cand[i]
	}
	return
}
