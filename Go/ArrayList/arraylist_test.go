package ArrayList

import "testing"

func TestFunc(t *testing.T)  {
	//nums := "25525511135"
	//nums := [][]int{{0,30},{5,10},{15,20}}
	nums := [][]int{{7,10},{2,4}}
	res := canAttendMeetings(nums)
	t.Log(res)
}