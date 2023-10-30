package leetcode

func NumberOfSteps(num int) int {
	steps := 0

	for num >= 1 {
		if num%2 == 0 {
			num /= 2
			steps++
			continue
		} else {
			num--
			steps++
		}
	}

	return steps
}
