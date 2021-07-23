package ArrayList

import "testing"

func TestRemoveDuplicates(t *testing.T)  {
	nums := []int{1,1}
	//nums := []int{1,2,3}
	res := removeDuplicates(nums)
	t.Log(res, nums[:res])
}
