package start

func runningSum(nums []int) []int {
	ans := make([]int, len(nums))

	ans[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		ans[i] = nums[i] + ans[i-1]
	}

	return ans
}
