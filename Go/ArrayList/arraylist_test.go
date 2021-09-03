package ArrayList

import "testing"

func TestFunc(t *testing.T)  {
	nums := []int{1,1}
	//nums := "25525511135"
	//nums := []int{1,2,3,4,5,6,7}
	//nums := [][]int{{1,2,3},{4,5,6},{7,8,9}}
	res := searchRange(nums, 1)
	t.Log(res)
}