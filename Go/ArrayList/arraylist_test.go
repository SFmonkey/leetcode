package ArrayList

import "testing"

func TestFunc(t *testing.T)  {
	//nums := "25525511135"
	//nums := [][]int{{0,30},{5,10},{15,20}}
	//nums := [][]int{{1,3},{2,6},{8,10},{15,18}}
	//nums := [][]int{{1,4},{2,3}}
	nums := []int{1,1,1,1,1,1,1,1}
	//res := dailyTemperatures(nums)
	//t.Log(res)
	minSubArrayLen(11, nums)
}

func TestSort(t *testing.T)  {
	nums := []int{4,5,2,3,1}
	sortArrayQuickSort(nums)
}