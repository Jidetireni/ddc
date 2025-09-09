package greedy

func MinOperations(nums []int) int {
	operations := 0
	for i := 1; i < len(nums); i++ {
		next := nums[i-1] + 1
		if next > nums[i] {
			operations += next - nums[i]
			nums[i] = next
		}
	}

	return operations
}
