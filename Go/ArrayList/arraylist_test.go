package ArrayList

import "testing"

func TestFunc(t *testing.T)  {
	//nums := []int{0,2,3,4,6,8,19}
	nums := "25525511135"
	//nums := []int{1,1,2,0,2,0}
	//nums := [][]int{{1,2,3},{4,5,6},{7,8,9}}
	res := restoreIpAddresses(nums)
	t.Log(res)
}