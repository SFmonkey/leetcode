package ArrayList

// 差分数组
func maximumPopulation(logs [][]int) int {
	population := make([]int, 101)			// 原数组
	diffPopulation := make([]int, 101)		// 差分数组
	// 这里不用初始化差分数组，因为原始数组全是0
	// 对于每一个子数组，相当于给每一个子数组区间内的元素+1
	for _, log := range logs {
		diffPopulation[log[0]-1950]++
		// 因为题目不包含右区间，这里-1950就行，否则-1949
		diffPopulation[log[1]-1950]--
	}
	// 还原数组
	maxIdx := 0								// 最大下标
	population[0] = diffPopulation[0]
	for i:=1; i<len(population); i++ {
		population[i] = population[i-1] + diffPopulation[i]
		if population[i] > population[maxIdx] {
			maxIdx = i
		}
	}
	return maxIdx+1950
}