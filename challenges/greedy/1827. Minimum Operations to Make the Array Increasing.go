package greedy

/*
	We are asked to return the number of operations needed to make the array
	strictly increasing, i.e., nums[i] < nums[i+1].

	My thought process was to act greedily: increase each element only when necessary,
	by at most +1 from the previous element. I calculate the required value (+1 of the
	previous element), then compare it with the current element.

	If the required value is greater, I find the number of operations by subtracting
	the current value from the required value. Then I update the current element to
	that required value.
*/

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
