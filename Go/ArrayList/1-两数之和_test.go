package ArrayList

import (
	"testing"
)

func TestTwoSum(t *testing.T)  {
	nums := []int{3,2,4}	// 可以这样声明array吗...(前面没有参数)
	target := 6
	res1 := twoSum1(nums, target)
	res2 := twoSum2(nums, target)
	t.Log(res1, res2)
}