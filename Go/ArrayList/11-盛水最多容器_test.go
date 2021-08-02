package ArrayList

import "testing"

func TestMaxArea(t *testing.T) {
	//in := []int{1,8,6,2,5,4,8,3,7}
	//in := []int{1,1}
	//in := []int{4,3,2,1,4}
	in := []int{1,2,1}
	res := maxArea(in)
	t.Log(res)
}
