package ArrayList

// 归并思想，从尾到头
func merge3( A []int ,  m int, B []int, n int )  {
	ap, bp := m-1, n-1
	idx := m+n-1
	for ap >= 0 && bp >= 0 {
		// 双指针比较，大的添加到尾部
		if A[ap] > B[bp] {
			A[idx] = A[ap]
			ap--
		} else {
			A[idx] = B[bp]
			bp--
		}
		idx--
	}
	// 把剩余的部分放到头部
	if ap < 0 {
		for i:=0; i<=bp; i++ {
			A[i] = B[i]
		}
	}
}
