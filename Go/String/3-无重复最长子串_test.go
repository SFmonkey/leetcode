package String

import "testing"

func TestLongSubStr(t *testing.T)  {
	s := "abcabcdbb"
	len_, res := lengthOfLongestSubstring(s)
	t.Log(len_, res)
	a := "abc"
	t.Log(string(a[1]))
}
