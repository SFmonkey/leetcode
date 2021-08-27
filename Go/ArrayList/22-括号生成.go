package ArrayList

import (
	"fmt"
	"strings"
)

// 回溯法, 从0开始
var input []string
var brackets []string
var bPath string
func generateParenthesis(n int) []string {
	input = []string{"(", ")"}
	brackets = []string{}
	bPath = ""
	backTBrackets(0, n)
	fmt.Println(brackets)
	return brackets
}

func backTBrackets(layer, n int)  {
	// 终止条件: 到达叶子节点
	if layer == 2*n {
		brackets = append(brackets, bPath)
		return
	}
	for i:=0; i<len(input); i++ {
		bPath += input[i]
		// 剪枝1: 左括号数大于n
		if strings.Count(bPath, "(") > n {
			bPath = bPath[:len(bPath)-1]
			continue
		}
		// 剪枝2: 右括号数大于做括号数
		if strings.Count(bPath, ")") > strings.Count(bPath, "(") {
			bPath = bPath[:len(bPath)-1]
			continue
		}
		backTBrackets(layer+1, n)
		bPath = bPath[:len(bPath)-1]
	}
}

