package ArrayList

import "testing"

func TestFindMedianSortedArrays(t *testing.T)  {
	nums1 := []int{1,3}
	nums2 := []int{2}
	mid := findMedianSortedArrays(nums1, nums2)
	t.Log(mid)
}
