package greedy

import "math"

func MinTimeToType(word string) int {
	cw := getClockWiseSteps(byte('a'), word[0])
	acw := getAntiClockWiseSteps(cw)
	secs := min(cw, acw) + 1

	for i := 0; i < len(word)-1; i++ {
		cw = getClockWiseSteps(word[i], word[i+1])
		acw = getAntiClockWiseSteps(cw)
		secs += min(cw, acw) + 1
	}
	return secs
}

func getClockWiseSteps(x, y byte) int {
	a := int(x) - int('a')
	b := int(y) - int('a')
	return (b - a + 26) % 26
}
func getAntiClockWiseSteps(cw int) int {
	return 26 - cw
}

// idomatic version
func minTimeToType(word string) int {
	time := steps('a', word[0]) + 1
	for i := 1; i < len(word); i++ {
		time += steps(word[i-1], word[i]) + 1
	}

	return time
}

func steps(a, b byte) int {
	diff := math.Abs(float64(a - b))
	return int(math.Min(diff, 26-diff))
}
