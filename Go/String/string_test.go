package String

import "testing"

func TestIsValid20(t *testing.T)  {
	res := isValid("()()[")
	t.Log(res)
}

func TestString(t *testing.T)  {
	res := reverseWords("the sky            is blue       ")
	t.Log(res)
}