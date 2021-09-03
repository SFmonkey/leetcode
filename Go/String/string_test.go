package String

import "testing"

func TestIsValid20(t *testing.T)  {
	res := isValid("()()[")
	t.Log(res)
}

func TestString(t *testing.T)  {
	res := multiply("401716832807512840963", "167141802233061013023557397451289113296441069")
	t.Log(res)
}