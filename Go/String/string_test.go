package String

import "testing"

func TestIsValid20(t *testing.T)  {
	res := isValid("()()[")
	t.Log(res)
}

func TestString(t *testing.T)  {
	res := addStrings("456", "77")
	t.Log(res)
}