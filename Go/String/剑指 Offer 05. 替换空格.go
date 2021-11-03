package String

func replaceSpace(s string) string {
	byteS := []byte(s)
	res := []byte{}
	for i:=0; i<len(byteS); i++ {
		if string(byteS[i]) == " " {
			res = append(res, '%')
			res = append(res, '2')
			res = append(res, '0')
		} else {
			res = append(res, byteS[i])
		}
	}
	return string(res)
}
