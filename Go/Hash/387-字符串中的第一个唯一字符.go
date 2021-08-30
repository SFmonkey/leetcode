package Hash

func firstUniqChar(s string) int {
	record := make(map[rune]int, len(s))
	for _, ch := range(s) {
		record[ch] += 1
	}
	for i, ch := range(s) {
		if record[ch] == 1 {
			return i
		}
	}
	return -1
}