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

// idomatic approach
func minOperations(nums []int) int {
	var ops int
	for i := 1; i < len(nums); i++ {
		if need := nums[i-1] + 1; need > nums[i] {
			ops += need - nums[i]
			nums[i] = need
		}
	}
	return ops
}
