package ArrayList

import "testing"

func TestFunc(t *testing.T)  {
	nums := []int{2, 3, 50, 75}
	//nums := "25525511135"
	//nums := []int{1,2,3,4,5,6,7}
	//nums := [][]int{{1,2,3},{4,5,6},{7,8,9}}
	findMissingRanges(nums, 0, 99)
}