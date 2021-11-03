package Hash

func canConstruct(ransomNote string, magazine string) bool {
	if magazine == "" {
		return false
	}
	record := make(map[byte]int)
	for i:=0; i<len(ransomNote); i++ {
		record[ransomNote[i]]++
	}
	for i:=0; i<len(magazine); i++ {
		num, ok := record[magazine[i]]
		if ok && num > 0 {
			record[magazine[i]]--
		}
	}
	for _, v := range record {
		if v != 0 {
			return false
		}
	}
	return true
}