package main

func stringShift(s string, shift [][]int) string {
	totalshift := 0

	for _, v := range shift {
		if v[0] > 0 {
			totalshift += v[1]
		} else {
			totalshift += v[1] * -1
		}
	}
	strlen := len(s)
	shiftmod := totalshift % strlen
	if shiftmod >= 0 {
		totalshift = shiftmod
	} else {
		totalshift = shiftmod + strlen
	}

	println("totalshift: ", totalshift)

	if totalshift == 0 {
		return s
	}
	head := s[strlen-totalshift : strlen]
	tail := s[0 : strlen-totalshift] //左闭右开（右侧不包含）

	return head + tail
}

func main() {
	var shift [][]int
	shift = [][]int{{0, 1}, {1, 2}}
	println("result: ", stringShift("abc", shift))
}
