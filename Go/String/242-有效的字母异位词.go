package String

func isAnagram(s string, t string) bool {
	record := make(map[rune]int, len(t))
	for _, ch := range t {
		record[ch] += 1
	}
	for _, ch := range s {
		if _, ok := record[ch]; ok {
			record[ch] -= 1
		} else {
			return false
		}
	}
	for _, v := range record {
		if v != 0 {
			return false
		}
	}
	return true
}