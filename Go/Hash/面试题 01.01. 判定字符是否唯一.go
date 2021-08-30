package Hash

func isUnique(astr string) bool {
	record := make(map[rune]bool, len(astr))
	for _, ch := range astr {
		if _, ok := record[ch]; !ok {
			record[ch] = true
		} else {
			return false
		}
	}
	return true
}