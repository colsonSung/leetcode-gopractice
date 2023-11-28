package leetcode

func confusingNumber(n int) bool {
	rotateArr := [10]int{0, 1, -1, -1, -1, -1, 9, -1, 8, 6} //用于识别翻转后的数字

	result := 0
	origin := n

	for n > 0 {
		x := n % 10
		n = n / 10
		if rotateArr[x] == -1 {
			return false
		}
		result = rotateArr[x] + result*10
	}

	println("origin %v", origin)
	println("result %v", result)

	return result != origin
}

//leetcode 1056.confusenum confusing number
//leetcode 1056.confusenum 旋转数字
