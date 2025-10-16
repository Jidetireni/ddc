package greedy

import (
	"math"
	"sort"
)

func MinMovesToSeat(seats []int, students []int) int {
	sort.Ints(seats)
	sort.Ints(students)

	count := 0
	for i := 0; i < len(seats); i++ {
		moves := seats[i] - students[i]
		count += int(math.Abs(float64(moves)))
	}

	return count
}

// idomatic version
func minMovesToSeat(seats []int, students []int) int {
	sort.Ints(seats)
	sort.Ints(students)

	moves := 0
	for i := range seats {
		moves += int(math.Abs(float64(seats[i] - students[i])))
	}

	return moves
}
