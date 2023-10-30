package leetcode

func MaximumWealth(accounts [][]int) int {
	max := 0
	sum := 0
	for i := 0; i < len(accounts); i++ {
		for j := 0; j < len(accounts[0]); j++ {
			sum += accounts[i][j]
		}
		if max < sum {
			max = sum
		}

		sum = 0
	}

	return max
}
