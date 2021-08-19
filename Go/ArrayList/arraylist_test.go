package ArrayList

import "testing"

func TestFunc(t *testing.T)  {
	//nums := []int{0,0,1,0,2,1,2,2}
	//nums := []int{1,1,2}
	//nums := []int{2,0,2,2,0}
	//nums := []int{1,1,2,0,2,0}
	//nums := [][]int{{1,2,3},{4,5,6},{7,8,9}}
	nums := [][]int{{5,1,9,11},{2,4,8,10},{13,3,6,7},{15,14,12,16}}
	rotate(nums)
}