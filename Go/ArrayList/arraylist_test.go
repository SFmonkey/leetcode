package ArrayList

import "testing"

func TestThreeSum15(t *testing.T)  {
	//in := []int{-1,0,1,2,-1,-4}
	in := []int{0}
	res := threeSum(in)
	t.Log(res)
}

func TestRemoveElement27(t *testing.T)  {
	in := []int{0,1,2,2,3,0,4,2}
	//in := []int{3,2,2,3}
	//in := []int{}
	t.Log(in)
	val := 2
	res := removeElement2(in, val)
	t.Log(in[:res])
}

func TestThreeSumClosest16(t *testing.T) {
	//in := []int{-1,2,1,-4}
	in := []int{1,1,1,0}
	// -4 -1 1 2
	target := -100
	res := threeSumClosest(in, target)
	t.Log(res)
}

func TestSearchInsert35(t *testing.T) {
	//in := []int{1,3,5,6}
	in := []int{1}
	target := 1
	res := searchInsert(in, target)
	t.Log(res)
}

func TestNextPermutation(t *testing.T)  {
	nums := []int{1,2,3,8,5,7,6,4}
	//nums := []int{3,2,1}
	nextPermutation(nums)
	t.Log(nums)
}