package ArrayList

import "testing"

func TestFunc(t *testing.T)  {
	nums := []int{0,0,1,0,2,1,2,2}
	//nums := []int{2,0,1}
	//nums := []int{2,0,2,2,0}
	//nums := []int{1,1,2,0,2,0}
	sortColors(nums)
	t.Log(nums)
}