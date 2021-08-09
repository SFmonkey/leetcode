package ArrayList

// 双指针+临时变量
func merge(nums1 []int, m int, nums2 []int, n int)  {
	if n == 0 {
		return
	}
	if m == 0 {
		nums1 = append(nums1, nums2...)
		return
	}
	if nums1[m-1] <= nums2[0] {
		for k:=m; k<m+n; k++ {
			nums1[k] = nums2[k-m]
		}
		return
	}
	var tmp, max int
	if nums1[m-1] > nums2[n-1] {
		tmp = nums1[m-1]
	} else {
		tmp = nums2[n-1]
	}
	max = tmp
	var i, j = 0, 0
	for i < m + n {
		if i >= m {
			nums1[i] = max
		}
		if j >= n {
			nums2[j] = max
		}
		if nums1[i] <= nums2[j] {
			if tmp < nums1[i] {
				nums1[i], tmp = tmp, nums1[i]
			}
			i++
		} else {
			if tmp < nums2[j] {
				nums1[i], tmp = tmp, nums1[i]
			} else {
				tmp = nums1[i]
				nums1[i] = nums2[j]
				j++
			}
			i++
		}
	}
}