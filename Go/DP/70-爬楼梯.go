package DP

/*
	f(10)=f(9)+f(8)
	f(9)=f(8)+f(7)
	f(8)=f(7)+f(6)
	...
	f(3)=f(2)+f(1)
	f(2)=2
	f(1)=1
	递归会超时，用带备忘录的递归可以
*/
func climbStairs2(n int) int {
	if 0 == n {
		return 1
	}
	if n <= 2 {
		return n
	}
	return climbStairs2(n-1) + climbStairs2(n-2)
}

func climbStairs(n int) int {
	if 0 == n {
		return 1
	}
	if n <= 2 {
		return n
	}
	a, b, tmp := 1, 2, 0
	for i:=3; i<=n; i++ {
		tmp = a + b
		a = b
		b = tmp
	}
	return tmp
}
