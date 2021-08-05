package ArrayList

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	if nums1[m-1] <= nums2[0] {
		if (m+n)%2 == 1 {return float64(nums2[0])} else {return float64(nums1[m-1]+nums2[0])/2}
	}
	if nums2[n-1] <= nums1[0] {
		if (m+n)%2 == 1 {return float64(nums1[0])} else {return float64(nums2[n-1]+nums1[0])/2}
	}
	var midLen, f int
	if 0 == (m+n)%2 {
		midLen = (m+n)/2-1
	} else {
		midLen = (m+n+1)/2-1
		f = 1
	}
	i, j := 0, 0
	for {
		if i + j == midLen {
			break
		}
		if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}

	}
	if 1 == f {
		fmt.Println("奇数", nums1[i], nums2[j], midLen)
		if nums1[i] > nums2[j] {return float64(nums1[i])} else {return float64(nums2[j])}
	}
	fmt.Println("偶数", i, j, midLen)
	return float64((nums1[i] + nums2[j]) / 2)
}
