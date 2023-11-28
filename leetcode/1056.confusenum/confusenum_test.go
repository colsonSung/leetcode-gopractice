package leetcode

import "testing"

func Test_confusingNumber(t *testing.T) {

	//confusingNumber(199)
	var cases = []struct {
		n      int
		output bool
	}{
		{6, true},
		{89, true},
		{11, false},
		{25, false},
		{916, false},
	}

	for _, c := range cases {
		x := confusingNumber(c.n)
		if x != c.output {
			t.Fail()
		}
	}
}
