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

<<<<<<< HEAD
func TestThreeSumClosest(t *testing.T)  {
	in := []int{-1,2,1,-4}
	// -4 -1 1 2
	target := 1
	res := threeSumClosest(in, target)
=======
func TestSearchInsert35(t *testing.T) {
	//in := []int{1,3,5,6}
	in := []int{1}
	target := 1
	res := searchInsert(in, target)
>>>>>>> 291125ee4a6c4edc0daea99c12856950af360363
	t.Log(res)
}
